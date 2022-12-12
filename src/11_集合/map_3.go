package main

import (
	"fmt"
)

func main() {
	m1 := map[string]string{"1":"1"}
	fmt.Println("update before：", m1)
	// update before： map[1:1]

	m1["1"] = "2"
	fmt.Println("update after:", m1)
	// update after: map[1:2]
}