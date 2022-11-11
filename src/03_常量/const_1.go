/**
 * 常量是一个简单的标识符，在程序允许时，不会被修改的量
 * 常量中的数据类型只可以是布尔型、数字型和字符串型
 * 格式：const identifier [type] = value
 * */

package main

import "fmt"

func main() {

	// 显类型定义
	const LENGTH int = 10

	// 隐式类型定义
	const WIDTH = 5

	var area int

	// 多重赋值
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d", area)
	println()
	println(a, b, c)
	// 常量值声明后可以不使用
	println(a, b)

}