package main

import (
	"fmt"
)

func printSlice(x , y []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(y), cap(y), y)
}

func main() {
	// 声明数组
	var arr = [5]int{0, 1, 2 ,3, 4}

	// 初始化切片
	slice_1 := arr[1:2]
	slice_2 := arr[1:4]

	printSlice(slice_1, slice_2)
	fmt.Println(arr)
	// len=1 cap=4 slice=[1]
	// len=3 cap=4 slice=[1 2 3]
	// [0 1 2 3 4]

	// 当追加的元素个数小于切片的剩余空间(cap-len)时，会直接影响原有的数组
	_ = append(slice_1, 1, 2)
	printSlice(slice_1, slice_2)
	fmt.Println(arr)
	// len=1 cap=4 slice=[1]
	// len=3 cap=4 slice=[1 1 2]
	// [0 1 1 2 4]

	// 会修改原来数组上的值
	_ = append(slice_1, 4)
	printSlice(slice_1, slice_2)
	fmt.Println(arr)
	// len=1 cap=4 slice=[1]
	// len=3 cap=4 slice=[1 4 2]
	// [0 1 4 2 4]

	// 重置数组，现在 slice_1 和 slice_2 都会改变
	arr = [5]int{0, 1, 2 ,3, 4}
	printSlice(slice_1, slice_2)
	fmt.Println(arr)
	// len=1 cap=4 slice=[1]
	// len=3 cap=4 slice=[1 2 3]
	// [0 1 2 3 4]
	
	// 现在添加的元素个数大于切片的剩余空间(cap-len)，此时将分配新的数组空间，不会影响原来的数组
	_ = append(slice_1, 1, 2, 3, 4)
	printSlice(slice_1, slice_2)
	fmt.Println(arr)
	// len=1 cap=4 slice=[1]
	// len=3 cap=4 slice=[1 2 3]
	// [0 1 2 3 4]
}