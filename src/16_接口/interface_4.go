package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName string
	bio string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    fmt.Println("Author: ", p.author.fullName())
    fmt.Println("Bio: ", p.author.bio)
}

func main() {
	author := author {"huang", "xiunian", "Java"}
	post := post{"Inheritance in Go", "Go Supports", author}
	post.details()
}