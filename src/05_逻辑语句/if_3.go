package main

import "fmt"

func main() {
	// 定义局部变量
	var a int = 100
	var b int = 200

	if a == 100 {
		// if 条件语句为 true 执行
		if b == 200 {
			// if 条件语句为 true 执行
			fmt.Println("a 的值为 100，b 的值为200")
		}
	}

	fmt.Println("a 的值为：", a)
	fmt.Println("b 的值为：", b)
}
