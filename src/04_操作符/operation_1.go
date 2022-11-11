/**
 * 算术运算符
 * */

package main

import "fmt"

func main() {
	a := 21
	b := 10
	var c int

	fmt.Println("a = ", a, "b = ", b)

	c = a + b
	fmt.Println("a + b = ", c)

	c = a - b
	fmt.Println("a - b = ", c)

	c = a * b
	fmt.Println("a * b = ", c)

	c = a / b
	fmt.Println("a / b = ", c)

	c = a % b
	fmt.Println("a % b = ", c)

	/**
	 * 和 Java 不一样，Go 的自增自减只能作为表达式使用，而不能用于赋值语句
	 * 
	 * b = a++   这是不允许的，会出现编译错误
	 * 
	 * */
	a++
	fmt.Println("a++ = ", a)

	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Println("a-- = ", a)
}