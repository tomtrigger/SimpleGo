package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
}

func (p person) fullName() {
	fmt.Printf("%s, %s", p.firstName, p.lastName)
}

func main() {
	p := person {
		firstName: "John",
		lastName: "smith",
	}

	defer p.fullName()
	fmt.Print("Welcome ")

}