package division

import (
	"fmt"
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	// 1. init
	InitData()

	// 2.
	fmt.Println("--------list---------")
	listProvince := ListProvince()
	fmt.Println(listProvince)
	// [{11    北京市} {12    天津市} {13    河北省}...]

	// 3. list next division
	nextDivision, err := ListNextDivision("5301")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(nextDivision)
	// [{53 5301 530101  市辖区} {53 5301 530102  五华区} {53 5301 530103  盘       龙区}..]

	// 4. get detail by code
	fmt.Println("---------detail--------")
	provinceDetail, err := GetDivisionDetail("53")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(provinceDetail)
	// [{53    云南省}]

	cityDetail, err := GetDivisionDetail("5301")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cityDetail)
	// [{53    云南省} {53 5301   昆明市}]

	countyDetail, err := GetDivisionDetail("530102")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(countyDetail)
	// [{53    云南省} {53 5301   昆明市} {53 5301 530102  五华区}]

}
