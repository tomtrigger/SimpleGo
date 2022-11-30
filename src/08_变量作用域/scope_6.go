package main 

import (
	"fmt"
)

func main() {
	// forTest1()
	forTest2()
}

func forTest1() {
	var a int = 0

	fmt.Println("for start")
	for a := 0; a < 10; a++ {
		fmt.Print(a, "\t")
	}
	fmt.Println("\nfor end")

	fmt.Println("a = ", a)

// for start
// 0       1       2       3       4       5       6       7       8       9
// for end
// a =  0
}

func forTest2() {
	var a int = 0

	fmt.Println("for start")
	// 注意这里是 a = 0,而不是 a := 0
	for a = 0; a < 10; a++ {
		fmt.Print(a, "\t")
	}
	fmt.Println("\nfor end")

	fmt.Println("a = ", a)

// for start
// 0       1       2       3       4       5       6       7       8       9
// for end
// a =  10
}