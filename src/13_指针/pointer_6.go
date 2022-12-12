package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
}

func main() {
	a := 59
	fmt.Println("value of a before function call is ", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is ", a)

	// value of a before function call is  59
	// value of a after function call is  55
}