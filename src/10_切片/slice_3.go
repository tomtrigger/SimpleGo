package main

import (
	"fmt"
)

func main() {
	// 创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	
	// 打印原始切片
	fmt.Println(numbers)
	// [0 1 2 3 4 5 6 7 8]

	// 打印子切片从索引1（包含）到索引4（不包含）
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	// numbers[1:4] == [1 2 3]

	// 默认上限为0
	fmt.Println("numbers[:3] ==", numbers[:3])
	// numbers[:3] == [0 1 2]

	// 默认下限为0
	fmt.Println("numbers[4:] ==", numbers[4:])
	// numbers[4:] == [4 5 6 7 8]

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
	// len=0 cap=5 slice=[]

	// 打印子切片从索引0（包含）到索引2（不包含）
	numbers2 := numbers[:2]
	printSlice(numbers2)
	// len=2 cap=9 slice=[0 1]

 	// 打印子切片从索引2（包含）到索引5（不包含）
	numbers3 := numbers[2:5]
	printSlice(numbers3)
	// len=3 cap=7 slice=[2 3 4]

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}