package main

import (
	"fmt"
)

func main() {
	// 定义局部变量
	var a int = 100
	var b int = 200

	fmt.Println("交换前的 a 值：", a)
	fmt.Println("交换前的 b 值：", b)

	// 通过函数交换值
	swap(a, b)

	fmt.Println("交换后的 a 值：", a)
	fmt.Println("交换后的 b 值：", b)
}

func swap(x, y int) (int, int) {
	var temp int

	temp = x
	x = y
	y = temp

	return x, y
}