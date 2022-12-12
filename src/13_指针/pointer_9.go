package main

import (
	"fmt"
)

func modify(sls []int) {
	sls[0] = 90
}

func main() {
	arr := [3]int{89, 90, 91}
	modify(arr[:])
	fmt.Println(arr)

	// [90 90 91]
}