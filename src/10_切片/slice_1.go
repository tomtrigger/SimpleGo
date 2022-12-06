package main

import (
	"fmt"
)

func main() {
	var numbers = make([]int, 6, 5)

	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice %v", len(x), cap(x), x)
}