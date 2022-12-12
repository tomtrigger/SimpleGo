package main

import "fmt"

func swap(x , y *int) {
	var temp int

	temp = *x  // 保存 x 地址的值
	*x = *y    // 将 y 赋值给 x
	*y = temp  // 将 temp 赋值给 y
}

func main() {
	var a int = 100
	var b int = 200

	fmt.Println("a = ", a, ", b = ", b)
	swap(&a, &b)

	fmt.Println("a = ", a, ", b = ", b)

	// a =  100 , b =  200
	// a =  200 , b =  100
}