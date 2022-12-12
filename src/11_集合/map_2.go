package main

import (
	"fmt"
)

func main() {
	var countryCapitalMap map[string]string

	// 创建集合
	countryCapitalMap = make(map[string]string)

	// map 插入 key-value 对，各个国家对应的首都
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	// 使用 key 输出 map 值
	for contry := range countryCapitalMap {
		fmt.Println("Capital of ", contry, " is ", countryCapitalMap[contry])
	}

	// 查看元素在集合中是否存在
	capital, ok := countryCapitalMap["United states"]

	if ok {
		fmt.Println("Capital of United states is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	// Capital of  France  is  Paris
	// Capital of  Italy  is  Rome
	// Capital of  Japan  is  Tokyo
	// Capital of  India  is  New Delhi
	// Capital of United States is not present

}	