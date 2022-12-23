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

	print(Book1)
	print(Book2)

	// Book title: Go 语言
	// Book author: www.runoon.com
	// Book subject: Go 语言教程
	// Book book_id: 123456
	// Book title: Python 教程
	// Book author: www.runoon.com
	// Book subject: Python 语言教程
	// Book book_id: 654321
}

func print(books Books) {
	fmt.Println("Book title:", books.title)
	fmt.Println("Book author:", books.author)
	fmt.Println("Book subject:", books.subject)
	fmt.Println("Book book_id:", books.book_id)
}
