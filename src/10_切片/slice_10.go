package main

import (
	"fmt"
)

func main() {
	slice_1 := []int{0, 1, 2, 3, 4}

	slice_2 := make([]int, 3)
	slice_3 := make([]int, 6)

	printSlice(slice_2, slice_3)
	// len=3 cap=3 slice=[0 0 0]
	// len=6 cap=6 slice=[0 0 0 0 0 0]

	// slice_2 的容量无法全部接收 slice_1 的数据，所以会丢弃元素 3 和 4
	copy(slice_2, slice_1)
	// slice_3 的容量大于 slice_1 的个数，所以会接收所有的元素
	copy(slice_3, slice_1)

	printSlice(slice_2, slice_3)
	// len=3 cap=3 slice=[0 1 2]
	// len=6 cap=6 slice=[0 1 2 3 4 0]
}

func printSlice(x, y []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(y), cap(y), y)
}