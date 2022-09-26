package division

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

var ProvinceList []Division
var CityList []Division
var CountyList []Division

func InitData(dir string) {
	ProvinceList = readCSV(dir + "/province.csv")
	CityList = readCSV(dir + "/city.csv")
	CountyList = readCSV(dir + "/county.csv")
}

func readCSV(fileName string) []Division {
	//准备读取文件
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	var list []Division
	// 不要第一行
	r.Read()
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
