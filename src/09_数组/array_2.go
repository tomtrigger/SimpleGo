package main

import (
	"fmt"
)

func main() {
	var i, j, k int

	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i = 0; i < 5; i++ {
		fmt.Print(balance[i], "\t")
	}

	fmt.Println()

	// 1000    2       3.4     7       50

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for j = 0; j < 5; j ++ {
		fmt.Print(balance2[j], "\t")
	}

	fmt.Println()

	// 1000    2       3.4     7       50

	balance3 := [...]float32{1:2.0, 50.0, 4:50.0}

	for k = 0; k < 5; k++ {
		fmt.Print(balance3[k], "\t")
	}

	// 0       2       50      0       50

}