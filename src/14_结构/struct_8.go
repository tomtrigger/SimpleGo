package main

import "computer"
import "fmt"

func main() {
	var spec computer.spec
	spec.Maker = "apple"
	spec.Price = 500
	fmt.Println("Spec: ", spec)
}