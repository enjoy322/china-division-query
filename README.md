### 简易的四级中国行政区划查询库

支持Redis缓存和普通缓存

#### 功能

> 注意：数据仅支持到街道/镇级

- 查询所有省份
- 查询某行政区信息
- 查询某行政区划下一级行政区
- 查询某行政区的详细政区信息
-

#### 使用

```shell
go get -u github.com/enjoy322/china-division-query@master
```

若需使用redis缓存,需使用以下库

```shell
github.com/go-redis/redis/v8
```

> 需要将行政区文件拷贝到你的项目启动根目录下,然后启动时指定文件目录
>

```go
package main

import (
	"fmt"
	division "github.com/enjoy322/china-division-query"
)

func main() {

	// 1. 使用redis作为缓存
	// client 为redis连接
	divisionClient := division.InitDivisionWithRedisAdapter("divisions", client)
	province := divisionClient.ListProvince()
	fmt.Println(province)
	// [{11    北京市} {12    天津市}...]

	//	2. 不适用redis作为缓存
	divisionClient := division.InitDivisionWithOutRedisAdapter("divisions")

	province, err := divisionClient.ListNextDivision(530102)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(province)
	// [{53 5301 530102 530102001 华山街道} {53 5301 530102 530102002 护国街道} ...
}

```
