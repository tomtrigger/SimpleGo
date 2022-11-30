package main

import (
	"fmt"
)

// 声明全局变量
var a int = 20

func main() {
	// 声明局部变量
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Println("main() 函数中 a = ", a)
	c = sum(a, b)
	fmt.Println("main() 函数中 c = ", c)

}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}
