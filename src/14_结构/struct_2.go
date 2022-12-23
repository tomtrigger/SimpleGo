package main

import (
	"fmt"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoon.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 123456

	Book2.title = "Python 教程"
	Book2.author = "www.runoon.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 654321

	fmt.Println("Book1 title:", Book1.title)
	fmt.Println("Bool1.author:", Book1.author)
	fmt.Println("Book1.subject:", Book1.subject)
	fmt.Println("Book1.book_id:", Book1.book_id)

	fmt.Println("Book2.title:", Book2.title)
	fmt.Println("Book2.author:", Book2.author)
	fmt.Println("Book2.subject:", Book2.subject)
	fmt.Println("Book2.book_id:", Book2.book_id)

	// Book1 title: Go 语言
	// Bool1.author: www.runoon.com
	// Book1.subject: Go 语言教程
	// Book1.book_id: 123456
	// Book2.title: Python 教程
	// Book2.author: www.runoon.com
	// Book2.subject: Python 语言教程
	// Book2.book_id: 654321
}