package main

import (
	"fmt"
	division "github.com/enjoy322/china-division-query"
	"log"
)

func main() {
	//1.
	client, err := RedisInit(Redis{
		Addr:     "127.0.0.1:6379",
		Db:       0,
		Password: "qwe123",
	})
	if err != nil {
		log.Fatalln(err)
	}
	_ = client

	limit := map[int]struct{}{
		53: struct{}{},
	}

	divisionClient := division.InitDivisionWithRedisAdapter("divisions", client, 3, limit)
	province := divisionClient.ListProvince()
	fmt.Println(province)

	// 2.
	//divisionClient := division.InitDivisionWithOutRedisAdapter("divisions")
	//
	//province, err := divisionClient.ListNextDivision(530102)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(province)
	// [{53 5301 530102 530102001 华山街道} {53 5301 530102 530102002 护国街道} ...
}
