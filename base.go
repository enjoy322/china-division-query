package division

type Division struct {
	ProvinceCode string `json:"province_code"`
	CityCode     string `json:"city_code"`
	CountyCode   string `json:"county_code"`
	TownCode     string `json:"town_code"`
	Name         string `json:"name"`
}
