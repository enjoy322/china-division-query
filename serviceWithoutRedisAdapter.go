package division

import "errors"

// ListProvince liat all provinces
func (s WithoutRedisAdapter) ListProvince() []Division {

	return provinces
}

// GetProvince get province information
func (s WithoutRedisAdapter) GetProvince(code int) (Division, error) {
	for _, division := range s.ListProvince() {
		if code == division.ProvinceCode {
			return division, nil
		}
	}
	return Division{}, errors.New("province code error")
}

// GetCity get city information
func (s WithoutRedisAdapter) GetCity(code int) (Division, error) {
	divisions := cities[code/100]
	for _, v := range divisions {
		if v.CityCode == code {
			return Division{
				ProvinceCode: code / 100,
				CityCode:     code,
				Name:         v.Name,
			}, nil
		}
	}

	return Division{}, errors.New("city code error")
}

// GetCounty get county information
func (s WithoutRedisAdapter) GetCounty(code int) (Division, error) {
	_ = counties
	divisions := counties[code/100]
	for _, v := range divisions {
		if v.CountyCode == code {
			return Division{
				ProvinceCode: code / 10000,
				CityCode:     code / 100,
				CountyCode:   code,
				Name:         v.Name,
			}, nil
		}
	}

	return Division{}, errors.New("county code error")
}

// GetTown get town information
func (s WithoutRedisAdapter) GetTown(code int) (Division, error) {
	divisions := towns[code/1000]
	for _, v := range divisions {
		if v.TownCode == code {
			return Division{
				ProvinceCode: code / 10000000,
				CityCode:     code / 100000,
				CountyCode:   code / 1000,
				TownCode:     code,
				Name:         v.Name,
			}, nil
		}
	}

	return Division{}, errors.New("town code error")
}

// GetDivisionDetail province-city-county,if the length of slice returned is 1
// it means code is of province
func (s WithoutRedisAdapter) GetDivisionDetail(code int) ([]Division, error) {
	switch numType(code) {
	case 2:
		division, err := s.GetProvince(code)
		if err != nil {
			return nil, err
		}
		return []Division{division}, nil
	case 4:
		division, err := s.GetCity(code)
		if err != nil {
			return nil, err
		}
		province, _ := s.GetProvince(division.ProvinceCode)
		return []Division{province, division}, nil
	case 6:
		division, err := s.GetCounty(code)
		if err != nil {
			return nil, err
		}
		province, _ := s.GetProvince(division.ProvinceCode)
		city, _ := s.GetCity(division.CityCode)
		return []Division{province, city, division}, nil
	case 9:
		division, err := s.GetTown(code)
		if err != nil {
			return nil, err
		}
		province, _ := s.GetProvince(division.ProvinceCode)
		city, _ := s.GetCity(division.CityCode)
		county, _ := s.GetCounty(division.CountyCode)
		return []Division{province, city, county, division}, nil
	default:
		return nil, errors.New("[error] division code error")
	}
}

// ListNextDivision list the next level divisions
func (s WithoutRedisAdapter) ListNextDivision(code int) ([]Division, error) {
	switch numType(code) {
	case 2:
		return s.listNextByProvince(code), nil
	case 4:
		return s.listNextByCity(code), nil
	case 6:
		return s.listNextByCounty(code), nil
	default:
		return nil, errors.New("code error")
	}
}

// ListNextByProvince list next level divisions by province code
func (s WithoutRedisAdapter) listNextByProvince(code int) []Division {
	return cities[code]
}

// ListNextByCity list next level divisions by city code
func (s WithoutRedisAdapter) listNextByCity(code int) []Division {
	return counties[code]
}

// list next level divisions by county code
func (s WithoutRedisAdapter) listNextByCounty(code int) []Division {
	return towns[code]
}
