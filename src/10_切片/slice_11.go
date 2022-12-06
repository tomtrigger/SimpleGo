package main

import (
	"fmt"
)

func main() {
	slice_1 := []int{1, 2, 3}
	slice_2 := []int{5, 6, 7, 8}

	copy(slice_2, slice_1)

	fmt.Println(slice_2)
	// [1 2 3 8]
}