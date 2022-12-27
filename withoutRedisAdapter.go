package division

type WithoutRedisAdapter struct {
}

func InitDivisionWithOutRedisAdapter(fileDir string) *WithoutRedisAdapter {
	dir = fileDir

	initDivision()
	return &WithoutRedisAdapter{}
}

func initDivision() {
	provinces = readProvince()

	cities = make(map[int][]Division)
	for _, v := range readCity() {
		cities[v.ProvinceCode] = append(cities[v.ProvinceCode], v)
	}

	counties = make(map[int][]Division)
	for _, v := range readCounty() {
		counties[v.CityCode] = append(counties[v.CityCode], v)
	}

	towns = make(map[int][]Division)
	for _, v := range readTown() {
		towns[v.CountyCode] = append(towns[v.CountyCode], v)
	}
}
