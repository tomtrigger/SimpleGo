package main

import "fmt"

func main() {
	var a int = 10

	for i := 0; i < 20; i++ {
		fmt.Println("a 的值为:", a)
		a++
		if a > 15 {
			// 使用 break 跳出循环
			break
		}
		
	}
}
