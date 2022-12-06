package main

import (
	"fmt"
)

func main() {
	var a = [5]int{1, 3, 4, 5, 6}

	for _, v := range a {
		fmt.Print(v, "\t")
	}
}