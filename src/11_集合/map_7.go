package main

import (
	"fmt"
)

func main() {
	m1 := map[string]int{"1":1}

	// m2 := map[string]int{"1":1}
	// fmt.Println(m1 == m2)

	fmt.Println(m1 == nil)
}