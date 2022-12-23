package main

import (
	"fmt"
)

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	map1[5] = 5.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Println("key: ", key, ", value : ", value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Println("key is: ", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Println("value is: ", value)
	}
}