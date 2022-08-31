package division

import "errors"

// ListProvince liat all provinces
func ListProvince() []Division {
	return ProvinceList
}

// ListCity all cities
func ListCity() []Division {
	return CityList
}

// ListCounty all counties
func ListCounty() []Division {
	return CountyList
}

// GetProvince get province information
func GetProvince(code string) (*Division, error) {
	for _, division := range ListProvince() {
		if code == division.ProvinceCode {
			return &division, nil
		}
	}
	return nil, errors.New("province code error")
}

// GetCity get city information
func GetCity(code string) (*Division, error) {
	for _, division := range ListCity() {
		if code == division.CityCode {
			return &division, nil
		}
	}
	return nil, errors.New("city code error")
}

// GetCounty get county information
func GetCounty(code string) (*Division, error) {
	for _, division := range ListCounty() {
		if code == division.CountyCode {
			return &division, nil
		}
	}
	return nil, errors.New("county code error")
}

// GetDivisionDetail province-city-county,if the length of slice returned is 1
// it means code is of province
func GetDivisionDetail(code string) ([]Division, error) {
	switch len(code) {
	case 2:
		division, err := GetProvince(code)
		if err != nil {
			return nil, err
		}
		return []Division{*division}, nil
	case 4:
		division, err := GetCity(code)
		if err != nil {
			return nil, err
		}
		province, _ := GetProvince(division.ProvinceCode)
		return []Division{*province, *division}, nil
	case 6:
		division, err := GetCounty(code)
		if err != nil {
			return nil, err
		}
		province, _ := GetProvince(division.ProvinceCode)
		city, _ := GetCity(division.CityCode)
		return []Division{*province, *city, *division}, nil
	default:
		return nil, errors.New("code error")
	}
}

// ListNextDivision list the next level divisions
func ListNextDivision(code string) ([]Division, error) {
	switch len(code) {
	case 2:
		return ListNextByProvince(code), nil
	case 4:
		return ListNextByCity(code), nil
	default:
		return nil, errors.New("code error")
	}
}

// ListNextByProvince list next level divisions by province code
func ListNextByProvince(code string) []Division {
	var list []Division
	for _, item := range CityList {
		if item.ProvinceCode == code {
			list = append(list, item)
		}
	}
	return list
}

// ListNextByCity list next level divisions by city code
func ListNextByCity(code string) []Division {
	var list []Division
	for _, item := range CountyList {
		if item.CityCode == code {
			list = append(list, item)
		}
	}
	return list
}
