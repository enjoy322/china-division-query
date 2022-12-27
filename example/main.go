package main

import (
	"fmt"
	division "github.com/enjoy322/china-division-query"
)

func main() {
	// 1.
	//client, err := RedisInit(Redis{
	//	Addr:     "127.0.0.1:6379",
	//	Db:       0,
	//	Password: "",
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//_ = client
	//
	//divisionClient := division.InitDivisionWithRedisAdapter("divisions", client)
	//province := divisionClient.ListProvince()
	//fmt.Println(province)

	// 2.
	divisionClient := division.InitDivisionWithOutRedisAdapter("divisions")

	province, err := divisionClient.ListNextDivision(530102)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(province)
	// [{53 5301 530102 530102001 华山街道} {53 5301 530102 530102002 护国街道} ...
}
