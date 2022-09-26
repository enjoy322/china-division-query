package division

import (
	"encoding/csv"
	"io"
	"log"
	"os"
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
			ProvinceCode: row[0],
			CityCode:     row[1],
			CountyCode:   row[2],
			TownCode:     row[3],
			Name:         row[4],
		})
	}
	return list
}
