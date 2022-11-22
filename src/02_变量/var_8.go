package main

import (
	"fmt"
)

func main() {
	// 值类型
	a := 3
	b := a

	fmt.Println(a, b)
	fmt.Println(&a, &b)

	fmt.Println("----------------")

	// 引用类型
	c := 4
	var d *int = &c

	fmt.Println(c, *d)
	fmt.Println(&c, d)

	c = 5

	fmt.Println(c, *d)
	fmt.Println(&c, d)

}
