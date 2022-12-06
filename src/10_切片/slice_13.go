package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	var b = a

	fmt.Println(a, b)

	b[0] = 4
	fmt.Println(a, b)
}