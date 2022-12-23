package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human // 匿名字段
	school string
}

type Employee struct {
	Human // 匿名字段
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I'm %s you can call me on %s \n", h.name, h.phone)
}

func main() {
	mark := Student{Human{"Mark", 22, "123456"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "67890"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}