package main

import (
	"fmt"
)

func main() {
	// 声明一个长度为 10 的数组
	var n [10] int

	var a, b int

	// 为数组 n 初始化元素
	for a = 0; a < 10; a++ {
		n[a] = a + 100
	}

	// 输出每个数组元素的值
	for b = 0; b < 10; b++ {
		fmt.Println(n[b])
	}
}