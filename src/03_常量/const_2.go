/**
 * 常量还可以用作枚举值
 * 
 * */
 
package main

import "unsafe"

/**
 * 常量的值必须是能够在编译时就能确定的，可以在其赋值表达式中涉及计算过程，但是所有用于计算的值
 * 必须在编译器就能获取，常量可以用len()、cap()、unsafe.Sizeof()函数计算表达式的值。
 * 
 * 正确  const a = 2 / 3
 * 错误  const b = getNumber()
 * */
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

 func main() {
 	println(a, b, c)
 }