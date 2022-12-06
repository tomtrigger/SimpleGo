package main

import (
	"fmt"
)

func printSlice(x, y []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(y), cap(y), y)
}

func main() {
	arr := [5]int{0, 1, 2, 3, 4}

	slice_1 := arr[:3]
	var slice_2 = append(slice_1, 4)
	printSlice(slice_1, slice_2)
	fmt.Println(arr)
	// len=3 cap=5 slice=[0 1 2]
	// len=4 cap=5 slice=[0 1 2 4]
	// [0 1 2 4 4]

	arr[0] = 11
	printSlice(slice_1, slice_2)
	fmt.Println(arr)
	// len=3 cap=5 slice=[11 1 2]
	// len=4 cap=5 slice=[11 1 2 4]
	// [11 1 2 4 4]
}