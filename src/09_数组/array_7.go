package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("数组 a :", a, " 长度是:", len(a))
	// 数组 a : [1 2 3 4 5]  长度是: 5

	b := [...]int{1:5, 10:9}
	fmt.Println("数组 b :", b, " 长度是:", len(b))
	// 数组 b : [0 5 0 0 0 0 0 0 0 0 9]  长度是: 11

}