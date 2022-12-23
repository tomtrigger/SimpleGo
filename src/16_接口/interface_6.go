package main

import "fmt"

type myint int
type mystr string

func main() {
	var i1 myint
	var i2 = 100

	fmt.Println(i1, i2)

	var name mystr
	name = "huang"
	var s1 string
	s1 = "xiunian"

	fmt.Println(name, s1)

	// fmt.Println(name == s1)
	// invalid operation: name == s1 (mismatched types mystr and string)
}