package main

import "fmt"

func main() {

	sum := 1

	// 无限循环
	for {
		sum += sum
	}

	fmt.Println(sum) // 无法输出

}
