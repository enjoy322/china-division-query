### 简易的中国行政区划查询库

#### 功能
>注意：数据仅支持到县区级
- 查询所有省份
- 查询所有城市
- 查询所有区县
- 查询某行政区划下一级行政区
- 查询某行政区的直接所有上级行政区
-

#### 使用

```shell
go get -u github.com/enjoy/china-division-query@master
```

> 需要将行政区文件拷贝到你的项目启动根目录下,然后启动时指定文件目录
> 
```go
package main

import (
	"fmt"
	division "github.com/enjoy/china-division-query"
)

func main() {
	// 初始化数据
	division.InitData("./files/")

	// 样例 
	list := division.ListProvince()
	fmt.Println(list)
	// [{11    北京市} {12    天津市}...]
}
```
