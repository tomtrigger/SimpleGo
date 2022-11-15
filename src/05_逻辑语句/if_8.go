package main

import "fmt"

func main() {
	if num := 9; num < 0 {
		fmt.Println(num, "是负数")
	} else if a := 1; num < a {
		fmt.Println(num, "值小于", a)
	} else {
		fmt.Println("num的值为:", num)
	}
}
