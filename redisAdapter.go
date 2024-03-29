package division

import (
	"context"
	"encoding/csv"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var provinces []Division
var cities map[int][]Division
var counties map[int][]Division
var towns map[int][]Division
var dir string

// cache
var cacheProvinceKey = "division_province"
var cacheCityKey = "division_city:"
var cacheCountyKey = "division_county:"
var cacheTownKey = "division_town:"

type RedisAdapter struct {
	rdb *redis.Client
}

func InitDivisionWithRedisAdapter(fileDir string, db *redis.Client, level int, limitProvince map[int]struct{}) *RedisAdapter {
	if db == nil {
		log.Fatalln("redis client is invalid")
	}
	dir = fileDir
	initWithRedis(db, level, limitProvince)
	return &RedisAdapter{
		rdb: db,
	}
}

func initWithRedis(rdb *redis.Client, level int, limitProvince map[int]struct{}) {
	log.Println("-------读取division------")
	start := time.Now()
	if level < 1 || level > 5 {
		log.Fatalln("level必须大于1且小于5")
	}
	if level > 0 {
		err := initProvince(readProvince(), rdb, limitProvince)
		if err != nil {
			return
		}
	}

	if level > 1 {
		err := initCity(readCity(), rdb, limitProvince)
		if err != nil {
			return
		}
	}
	if level > 2 {
		err := initCounty(readCounty(), rdb, limitProvince)
		if err != nil {
			return
		}
	}
	if level > 3 {
		err := initTown(readTown(), rdb, limitProvince)
		if err != nil {
			return
		}
	}

	log.Println("---------division 写入redis完成：", time.Now().Sub(start).Seconds(), "------------")
}

func initTown(list []Division, rdb *redis.Client, limitProvince map[int]struct{}) error {
	pip := rdb.Pipeline()
	ctx := context.Background()
	for _, v := range list {
		if _, ok := limitProvince[v.ProvinceCode]; ok {
			pip.ZAdd(ctx, cacheTownKey+strconv.Itoa(v.CountyCode), &redis.Z{Score: float64(v.TownCode), Member: v.Name})
		}
	}
	_, err := pip.Exec(ctx)
	return err
}

func initCounty(list []Division, rdb *redis.Client, limitProvince map[int]struct{}) error {
	pip := rdb.Pipeline()
	ctx := context.Background()
	for _, v := range list {
		if _, ok := limitProvince[v.ProvinceCode]; ok {
			pip.ZAdd(ctx, cacheCountyKey+strconv.Itoa(v.CityCode), &redis.Z{Score: float64(v.CountyCode), Member: v.Name})
		}
	}
	_, err := pip.Exec(ctx)
	return err
}

func initCity(list []Division, rdb *redis.Client, limitProvince map[int]struct{}) error {
	pip := rdb.Pipeline()
	ctx := context.Background()
	for _, v := range list {
		if _, ok := limitProvince[v.ProvinceCode]; ok {
			pip.ZAdd(ctx, cacheCityKey+strconv.Itoa(v.ProvinceCode), &redis.Z{Score: float64(v.CityCode), Member: v.Name})
		}
	}
	_, err := pip.Exec(ctx)
	return err
}

func initProvince(list []Division, rdb *redis.Client, limitProvince map[int]struct{}) error {
	pip := rdb.Pipeline()
	ctx := context.Background()
	for _, v := range list {
		if _, ok := limitProvince[v.ProvinceCode]; ok {
			pip.ZAdd(ctx, cacheProvinceKey, &redis.Z{Score: float64(v.ProvinceCode), Member: v.Name})
		}
	}
	_, err := pip.Exec(ctx)
	return err
}

func readProvince() []Division {
	return readCSV(dir + "/province.csv")
}
func readCity() []Division {
	return readCSV(dir + "/city.csv")
}
func readCounty() []Division {
	return readCSV(dir + "/county.csv")
}
func readTown() []Division {
	return readCSV(dir + "/town.csv")
}

func readCSV(fileName string) []Division {
	//准备读取文件
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer func(fs *os.File) {
		err := fs.Close()
		if err != nil {

		}
	}(fs)
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	var list []Division
	// 不要第一行
	_, _ = r.Read()
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		list = append(list, Division{
			ProvinceCode: backInt(row[0]),
			CityCode:     backInt(row[1]),
			CountyCode:   backInt(row[2]),
			TownCode:     backInt(row[3]),
			Name:         row[4],
		})
	}
	return list
}

func backInt(codeStr string) int {
	if len(codeStr) == 0 {
		return 0
	}
	i, err := strconv.Atoi(codeStr)
	if err != nil {
		log.Fatalln("[error] 数据有误. ", err)
	}
	return i
}
