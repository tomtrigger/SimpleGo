package main

import (
	"fmt"
)

func main() {
	b := 255
	a := &b

	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)

	*a++

	fmt.Println("value of b is", *a)

	// address of b is 0xc00000e0a8
	// value of b is 255
	// value of b is 256
}