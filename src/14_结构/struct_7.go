package main

import (
	"fmt"
)

type Address struct {
	city, state string
}

type Person struct {
	name string
	age int
	Address
}

func main() {
	var p Person
	p.name = "Naveen"
	p.age = 50
	p.Address = Address {
		city: "Chicago",
		state: "Illinois",
	}

	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("City: ", p.city)   // city 是提升字段
	fmt.Println("State: ", p.state) // state 是提升字段
}