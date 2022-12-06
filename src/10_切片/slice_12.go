package main

import (
	"fmt"
)

func main() {
	slice_1 := []int{1, 2, 3}
	slice_2 := make([]int, 3)
	copy(slice_2, slice_1)
	fmt.Println(slice_1, slice_2)
	// [1 2 3] [1 2 3]

	slice_2[0] = 5
	fmt.Println(slice_1, slice_2)
	// [1 2 3] [5 2 3]
}