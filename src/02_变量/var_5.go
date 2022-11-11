package main

import "fmt"

func main() {
	/*
		如果变量已经使用 var 声明过了，在使用 := 声明变量，就产生编译错误，如下：
		var a int
		a = 1

		a:=1 // 错误

	*/
	f := "huangxiunian" // var f string = "huangxiunian"

	fmt.Println(f)
}