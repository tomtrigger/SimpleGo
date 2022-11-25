package main

import (
	"fmt"
)

func main() {
	a, b := swap("huang", "xiunian")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}