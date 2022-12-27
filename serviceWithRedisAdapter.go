package division

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// ListProvince liat all provinces
func (s RedisAdapter) ListProvince() []Division {
	zCard := s.rdb.ZCard(context.Background(), cacheProvinceKey)
	val := s.rdb.ZRangeWithScores(context.Background(), cacheProvinceKey, 0, zCard.Val()).Val()
	var list []Division
	for _, z := range val {
		list = append(list, Division{
			ProvinceCode: int(z.Score),
			Name:         z.Member.(string),
		})
	}
	return list
}

// GetProvince get province information
func (s RedisAdapter) GetProvince(code int) (Division, error) {
	i := strconv.Itoa(code)
	val := s.rdb.ZRangeByScore(context.Background(), cacheProvinceKey, &redis.ZRangeBy{Min: i, Max: i}).Val()
	if len(val) < 1 {
		return Division{}, errors.New("province code error")
	}
	return Division{ProvinceCode: code, Name: val[0]}, nil
}

// GetCity get city information
func (s RedisAdapter) GetCity(code int) (Division, error) {
	i := strconv.Itoa(code)
	err := errors.New("city code error")
	if len(i) != 4 {
		return Division{}, err
	}
	val := s.rdb.ZRangeByScore(context.Background(), cacheCityKey+i[:2], &redis.ZRangeBy{Min: i, Max: i}).Val()
	if len(val) < 1 {
		return Division{}, err
	}

	return Division{ProvinceCode: code / 100, CityCode: code, Name: val[0]}, nil
}

func (s RedisAdapter) GetCounty(code int) (Division, error) {
	i := strconv.Itoa(code)
	err := errors.New("county code error")
	if len(i) != 6 {
		return Division{}, err
	}
	val := s.rdb.ZRangeByScore(context.Background(), cacheCountyKey+i[:4], &redis.ZRangeBy{Min: i, Max: i}).Val()
	if len(val) < 1 {
		return Division{}, err
	}

	return Division{ProvinceCode: code / 10000, CityCode: code / 100, CountyCode: code, Name: val[0]}, nil
}

func (s RedisAdapter) GetTown(code int) (Division, error) {
	i := strconv.Itoa(code)
	err := errors.New("county code error")
	if len(i) != 9 {
		return Division{}, err
	}
	val := s.rdb.ZRangeByScore(context.Background(), cacheTownKey+i[:6], &redis.ZRangeBy{Min: i, Max: i}).Val()
	if len(val) < 1 {
		return Division{}, err
	}
	return Division{ProvinceCode: code / 10000000, CityCode: code / 100000, CountyCode: code / 1000, TownCode: code, Name: val[0]}, nil
}

// GetDivisionDetail province-city-county,if the length of slice returned is 1
// it means code is of province
func (s RedisAdapter) GetDivisionDetail(code int) ([]Division, error) {
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
func (s RedisAdapter) ListNextDivision(code int) ([]Division, error) {
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
func (s RedisAdapter) listNextByProvince(code int) []Division {
	i := strconv.Itoa(code)
	key := cacheCityKey + i
	zCard := s.rdb.ZCard(context.Background(), key)
	val := s.rdb.ZRangeWithScores(context.Background(), key, 0, zCard.Val()).Val()
	var list []Division
	for _, z := range val {
		list = append(list, Division{
			ProvinceCode: code,
			CityCode:     int(z.Score),
			Name:         z.Member.(string),
		})
	}
	return list
}

// ListNextByCity list next level divisions by city code
func (s RedisAdapter) listNextByCity(code int) []Division {
	i := strconv.Itoa(code)
	key := cacheCountyKey + i
	zCard := s.rdb.ZCard(context.Background(), key)
	val := s.rdb.ZRangeWithScores(context.Background(), key, 0, zCard.Val()).Val()
	var list []Division
	for _, z := range val {
		list = append(list, Division{
			ProvinceCode: code / 100,
			CityCode:     code,
			CountyCode:   int(z.Score),
			Name:         z.Member.(string),
		})
	}
	return list
}

// listNextByCounty list next level divisions by county code
func (s RedisAdapter) listNextByCounty(code int) []Division {
	i := strconv.Itoa(code)
	key := cacheTownKey + i
	zCard := s.rdb.ZCard(context.Background(), key)
	val := s.rdb.ZRangeWithScores(context.Background(), key, 0, zCard.Val()).Val()
	var list []Division
	for _, z := range val {
		list = append(list, Division{
			ProvinceCode: code / 10000,
			CityCode:     code / 100,
			CountyCode:   code,
			TownCode:     int(z.Score),
			Name:         z.Member.(string),
		})
	}
	return list
}

func numType(code int) int {
	i := 1
	for code > 9 {
		code = code / 10
		i++
	}
	return i
}
