package main

import (
	"fmt"
)

func main() {
	const (
		a = 1
		b
		c
		d
	)

	fmt.Println(a)
	// b、c、d 没有初始化，使用上一行（即 a）的值
	fmt.Println(b) // 输出1
	fmt.Println(c) // 输出1
	fmt.Println(d) // 输出1
}