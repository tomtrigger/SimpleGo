package main

import (
	"fmt"
)

func main() {
	var a int = 10
	fmt.Println("变量的地址:", &a)
	// 变量的地址: 0xc0000a6058
}