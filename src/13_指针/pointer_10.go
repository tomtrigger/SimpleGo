package main 

import "fmt"

func main() {
	var a int
	var ptr *int
	var pttr **int

	a = 3000

	// 指针 ptr 地址
	ptr = &a

	// 指向指针 ptr 地址 
	pttr = &ptr

	fmt.Println("变量 a = ", a)
	fmt.Println("指针变量 *ptr = ", *ptr)
	fmt.Println("指向指针的指针变量 **pttr = ", **pttr)

	// 变量 a =  3000
	// 指针变量 *ptr =  3000
	// 指向指针的指针变量 **pttr =  3000
}