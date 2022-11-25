package main

import (
	"fmt"
)

func main() {
	var a int = 100
	var b int = 200

	fmt.Println("交换前 a 的值：", a)
	fmt.Println("交换前 b 的值：", b)

	/*
	 * 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	swap(&a, &b)

	fmt.Println("交换后 a 的值:", a)
	fmt.Println("交换后 b 的值:", b)
}

func swap(x *int, y *int) {
	var temp int

	temp = *x
	*x = *y
	*y = temp
}