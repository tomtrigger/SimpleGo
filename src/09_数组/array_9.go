package main

import (
	"fmt"
)

func main() {
	var a = [4]int{1, 2, 3, 4}
	b := a
	b[0] = 10

	fmt.Println(a) // [1 2 3 4]
	fmt.Println(b) // [10 2 3 4]
}