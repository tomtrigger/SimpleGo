package main

import (
	"fmt"
)

func main() {
	var arr = [3]int{0, 1, 2}
	slice := arr[:1]
	changeSlice(slice)
	fmt.Println(arr)

	slice = arr[:2]
	fmt.Println(slice)
}

func changeSlice(slice []int) {
	_ =append(slice, 100)
}