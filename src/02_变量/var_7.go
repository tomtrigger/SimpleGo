package main

import "fmt"

// 全局变量允许声明但是不使用
var b, c , d int

func main() {
	// 如果定义了局部变量而不使用，在编译的时候就会报错
	var a string = "abc"
	fmt.Println("hello world", a)

	// 多变量可以在同一行进行赋值
	var aa, bb int
	var cc string

	aa, bb, cc = 1, 2, "aa"
	fmt.Println(aa, bb, cc)

	// 如果变量未声明，可以想下面那样使用
	aaa, bbb, ccc := 4, 5, "555"
	fmt.Println(aaa, bbb, ccc)

	// 交换2个变量的值，前提是2个变量的数据类型要一致
	fmt.Println("------")
	fmt.Println(aaa, bbb)
	aaa, bbb = bbb, aaa
	fmt.Println(aaa, bbb)

	/**
	 * 空白标识符 _被用于抛弃值，如值5在： _, b = 5, 7 中被抛弃
	 * 
	 * (_)实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用
	 * 所有被声明的变量，但有时你并不需要使用从一个函数得到的所有的返回值。
	 * 
	 * 并行赋值也被用于当一个函数返回多个返回值时，比如这里的val 和错误 err 是通过
	 * 调用 Func1 函数同时得到： val, err = Func1(var1).
	 * */
	 // 只获取函数返回值的后两个
	 _, num, str := numbers()
	 fmt.Println(num, str)

}

// 一个可以返回多个值的函数
func numbers()(int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}