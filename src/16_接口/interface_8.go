package main

import (
	"fmt"
)

type myint int
type myint2 = int // 不是重新定义，只是给 int 起别名

func main() {
	var i1 myint
	var i2 = 100
	// i1 = i2 // cannot use i2 (variable of type int) as type myint in assignment
	// fmt.Println(i1 == i2) // invalid operation: i1 == i2 (mismatched types myint and int)

	var i3 myint2
	i3 = i2
	fmt.Println(i1, i2, i3)
}