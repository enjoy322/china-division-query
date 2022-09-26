package division

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FuzzGetCity(f *testing.F) {
	InitData("./files/")

	f.Add(5301)
	f.Fuzz(func(t *testing.T, code int) {
		info, err := GetCity(code)
		if err != nil {
			t.Error(err)
		}
		t.Log(info)
		// {53 5301   昆明市}
		assert.Equal(t, "昆明市", info.Name)
	})
}

func FuzzGetCounty(f *testing.F) {
	InitData("./files/")

	f.Add(530102)
	f.Fuzz(func(t *testing.T, code int) {
		info, err := GetCounty(code)
		if err != nil {
			t.Error(err)
		}
		t.Log(info)
		// {53 5301 530102  五华区}
		assert.Equal(t, 5301, info.CityCode)
		assert.Equal(t, 530102, info.CountyCode)
		assert.Equal(t, "五华区", info.Name)
	})
}

func FuzzGetDivisionDetail(f *testing.F) {
	InitData("./files/")

	f.Add(53)
	f.Add(5301)
	f.Add(530102)
	f.Fuzz(func(t *testing.T, code int) {
		info, err := GetDivisionDetail(code)
		if err != nil {
			t.Error(err)
		}
		t.Log(info)
		// [{53    云南省}]
		// [{53    云南省} {53 5301      昆明市}]
		// [{53    云南省} {53 5301      昆明市} {53 5301 530102  五华区}]
	})
}

func FuzzGetProvince(f *testing.F) {
	InitData("./files/")

	f.Add(53)
	f.Fuzz(func(t *testing.T, code int) {
		info, err := GetProvince(code)
		if err != nil {
			t.Error(err)
		}
		t.Log(info)
		// {53    云南省}
		assert.Equal(t, "云南省", info.Name)
	})
}

func FuzzListNextByCity(f *testing.F) {
	InitData("./files/")

	f.Add(5301)
	f.Fuzz(func(t *testing.T, code int) {
		byCity := ListNextByCity(code)
		t.Log(byCity)
		// [{53 5301 530101  市辖区} {53    5301 530102  五华区} ... ]
	})
}

func FuzzListNextByProvince(f *testing.F) {
	InitData("./files/")

	f.Add(53)
	f.Fuzz(func(t *testing.T, code int) {
		d := ListNextByProvince(code)
		t.Log(d)
		// [{53 5301   昆明市} {53 5   303   曲靖市} ...]
	})
}

func FuzzListNextDivision(f *testing.F) {
	InitData("./files/")

	f.Add(5301)
	f.Fuzz(func(t *testing.T, code int) {
		divisions, err := ListNextDivision(code)
		if err != nil {
			t.Error(err)
		}
		t.Log(divisions)
		// [{53 5301 530101  市辖区}    {53 5301 530102  五华区} ... ]
	})
}

func FuzzListProvince(f *testing.F) {
	InitData("./files/")

	f.Add(1)
	f.Fuzz(func(t *testing.T, i int) {
		province := ListProvince()
		assert.Equal(t, 31, len(province))
	})
}
