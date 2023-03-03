package division

import "log"

type WithoutRedisAdapter struct {
}

func InitDivisionWithOutRedisAdapter(fileDir string, level int) *WithoutRedisAdapter {
	dir = fileDir

	initDivision(level)
	return &WithoutRedisAdapter{}
}

func initDivision(level int) {
	if level < 1 || level > 5 {
		log.Fatalln("level必须大于1且小于5")
	}
	provinces = readProvince()
	cities = make(map[int][]Division)

	counties = make(map[int][]Division)

	towns = make(map[int][]Division)

	if level > 1 {
		for _, v := range readCity() {
			cities[v.ProvinceCode] = append(cities[v.ProvinceCode], v)
		}
	}
	if level > 2 {
		for _, v := range readCounty() {
			counties[v.CityCode] = append(counties[v.CityCode], v)
		}
	}
	if level > 3 {
		for _, v := range readTown() {
			towns[v.CountyCode] = append(towns[v.CountyCode], v)
		}
	}

}
