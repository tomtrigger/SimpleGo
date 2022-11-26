package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2

	defer fmt.Println(a)

	fmt.Println(b)
}