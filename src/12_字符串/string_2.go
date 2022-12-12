package main

import (
	"fmt"
)

func main() {
	s := "hello world"

	for i:=0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}

	fmt.Println("")

	for i:=0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}