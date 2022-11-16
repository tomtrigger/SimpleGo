package main

import "fmt"

func main() {
	// 定义局部变量
	var a int = 5

LOOP:
	for a < 20 {
		fmt.Println("a 的值为:", a)
		a++
		// 跳过本次迭代
		if a == 15 {
			a++
			goto LOOP
		}
	}
}
