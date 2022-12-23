package main

import "fmt"

type Student struct {

}

func main() {
	// var i1 interface{} = new(Student)
	// s1 := i1.(Student)  // 不安全，如果断言失败，会直接 panic
	// fmt.Println(s1)

	var i2 interface{} = new(Student)
	s2, ok := i2.(Student) // 安全，断言失败也不会 panic，只是 ok 的值为 false
	if ok {
		fmt.Println(s2)
	}

}