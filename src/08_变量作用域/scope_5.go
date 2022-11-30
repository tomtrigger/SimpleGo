package main

import (
	"fmt"
)

func main() {
	// test1()
	// test2()
	test3()
}

func test1() {
	a := 5
	{
		a := 3
		fmt.Println("in a = ", a)
	}
	fmt.Println("out a = ", a)

// in a =  3
// out a =  5
}

func test2() {
	a := 5
	{
		fmt.Println("in a = ", a)
	}
	fmt.Println("out a = ", a)

// in a =  5
// out a =  5
}

func test3() {
	a := 5
	{
		a := 3
		fmt.Println("a = ", a)
	}

// a declared but not used
}