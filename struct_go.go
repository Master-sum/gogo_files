package main

import "fmt"

type Book struct {
	title  string
	author string
	id     int
}

func main() {
	var book Book
	book.title = "红楼梦"
	book.author = "曹雪芹"
	book.id = 9
	fmt.Print(book)
}
