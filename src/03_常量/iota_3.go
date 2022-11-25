package main

import (
	"fmt"
)

func main () {
	// 使用 iota 的常量的类型只能使用整型或者浮点型
	const a float32 = iota

	fmt.Println(a)
}