package main

import (
	"fmt"
)

func main() {
	var numbers []int

	printSlice(numbers)

	if numbers == nil {
		fmt.Println("切片是 nil")
	}

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// len=0 cap=0 slice=[]
// 切片是 nil