package main

import (
	"fmt"
)

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("Person -->", p.name)
}

// 类型别名
type People = Person

type Student struct {
	// 嵌入两个结构
	Person
	People
}

func (p People) show2() {
	fmt.Println("People -->", p.name)
}

func main() {
	var s Student

	// s.name = "haung" // ambiguous selector s.name
	s.Person.name = "huang"
	s.People.name = "xiunian"

	// s.show() // ambiguous selector s.show
	s.Person.show()
	s.Person.show2()
	s.People.show()
	s.People.show2()
}