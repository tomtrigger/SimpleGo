package main

import (
	"fmt"
)

func modiy(arr *[3]int) {
	arr[0] = 90
}

func main() {
	arr := [3]int{89, 90, 91}
	modiy(&arr)
	fmt.Println(arr)

	// [90 90 91]
}