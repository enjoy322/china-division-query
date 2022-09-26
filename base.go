package division

type Division struct {
	ProvinceCode int    `json:"province_code"`
	CityCode     int    `json:"city_code"`
	CountyCode   int    `json:"county_code"`
	TownCode     int    `json:"town_code"`
	Name         string `json:"name"`
}
