/**
 * 运算符优先级
 * 
 * 有些运算符拥有较高的优先级，二元运算符的运算方向均是从左到右。当然也可以通过使用括号来临时
 * 提升某个表达式的整体运算优先级
 * 
 * 优先级 			运算符
 * 
 * 5				* / % << >> &
 * 4 				+ - | ^
 * 3 				== != < <= > >=
 * 2 				&&
 * 1 				||
 * */

 package main

 import "fmt"

 func main() {
 	var a int = 20
 	var b int = 10
 	var c int = 15
 	var d int = 5

 	fmt.Println("(a + b) * c / d = ", (a + b) * c /d)
 	fmt.Println("((a + b) * c) / d = ", ((a + b) * c) / d)
 	fmt.Println("(a + b) * (c / d) = ",(a + b) * (c / d))
 	fmt.Println("a + (b * c) / d = ", a + (b * c) / d)
 }