package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2

	c := max(a, b)
	fmt.Println("最大的值为:", c)
}

func max(num1, num2 int) int {
	// 声明局部变量
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}