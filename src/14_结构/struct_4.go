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

	printBook(&Book1)
	printBook(&Book2)

	// book title:  Go 语言
	// book author:  www.runoon.com
	// book subject:  Go 语言教程
	// book book_id:  123456
	// book title:  Python 教程
	// book author:  www.runoon.com
	// book subject:  Python 语言教程
	// book book_id:  654321
}

func printBook(book *Books) {
	fmt.Println("book title: ", book.title)
	fmt.Println("book author: ", book.author)
	fmt.Println("book subject: ", book.subject)
	fmt.Println("book book_id: ", book.book_id)
}