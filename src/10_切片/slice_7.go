package main

import (
	"fmt"
)

func main() {
	var arr1 = []int{1, 2, 3}
	var arr2 = []int{4, 5, 6}
	var arr3 = []int{7, 8, 9}

	var s1 = append(append(arr1, arr2...), arr3...)
	fmt.Println(s1)

	// [1 2 3 4 5 6 7 8 9]

}