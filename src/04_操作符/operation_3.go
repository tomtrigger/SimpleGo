/**
 * 逻辑运算符
 * */
package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false


	fmt.Println("a && b:", a && b)
	fmt.Println("a || b:", a || b)
	fmt.Println("!(a && b):", !(a && b))
}