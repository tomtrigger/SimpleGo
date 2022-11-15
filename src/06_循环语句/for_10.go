package main

import "fmt"

func main() {

	// 不适用标记
	fmt.Println("----continue----")

	for i := 0; i < 3; i++ {
		fmt.Println("i 的值：", i)
		for j := 0; j < 5; j++ {
			fmt.Println("j 的值：", j)
			continue
		}
	}

	// 使用标记
	fmt.Println("----continue label----")

re:
	for i := 0; i < 3; i++ {
		fmt.Println("i 的值：", i)
		for j := 0; j < 5; j++ {
			fmt.Println("j 的值：", j)
			continue re
		}
	}

}
