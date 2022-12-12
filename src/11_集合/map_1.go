package main

import (
	"fmt"
)

func main() {
	var m1 = map[string]string{"1":"1"}
	m1["2"] = "2"

	m2 := make(map[string]float32)
	m2["Go"] = 5.0

	fmt.Println(m1, m2)
}