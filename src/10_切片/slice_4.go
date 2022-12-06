package main

import (
	"fmt"
)

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {
	var numbers []int
	printSlice(numbers)
	// len=0 cap=0 slice=[]

	// 允许追加空切片
	numbers = append(numbers, 0)
	printSlice(numbers)
	// len=1 cap=1 slice=[0]

	// 向切片中添加一个元素
	numbers = append(numbers, 1)
	printSlice(numbers)
	// len=2 cap=2 slice=[0 1]

	// 同时添加多个元素
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
	// len=5 cap=6 slice=[0 1 2 3 4]

	// 创建切片 numbers1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), cap(numbers) * 2)
	printSlice(numbers1)
	// len=5 cap=12 slice=[0 0 0 0 0]

	// 拷贝 numbers 的内容到 numbers1
	copy(numbers1, numbers)
	printSlice(numbers1)
	// len=5 cap=12 slice=[0 1 2 3 4]

}