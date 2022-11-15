package main

import "fmt"

func main() {
	// 局部变量定义
	var a int = 20

	// 判断布尔表达式
	if a < 20 {
		// 如果条件为 true 则执行以下语句
		fmt.Println("a 小于 20")
	} else {
		// 如果条件为 false 则执行以下语句
		fmt.Println("a 大于 20")
	}

	fmt.Println("a 的值为：", a)
}