package main

import "fmt"

func main() {
	/*
		算术运算符：+，-，*, /, %, ++, --
		+
		-
		*：乘法
		/：取商，两个数相除，取商
		%：取余，取模，两个数相除，取余数

		++：自增1
			i++
		--：自减1
			i--
	*/

	// 整数
	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d + %d = %d \n", a, b, sum)

	sub := a - b
	fmt.Printf("%d - %d = %d \n", a, b, sub)

	mul := a * b
	fmt.Printf("%d * %d = %d \n", a, b, mul)

	div := a / b // 取商
	mod := a % b // 取余
	fmt.Printf("%d / %d = %d \n", a, b, div)
	fmt.Printf("%d %% %d = %d \n", a, b, mod)

	c := 3
	c++ // 自增1
	fmt.Println(c)

	c-- // 自减1
	fmt.Println(c)

}
