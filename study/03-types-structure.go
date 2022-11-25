package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
}

func main() {
	var Book1 Books
	Book1.title = "Book 1 Title"
	Book1.author = "Book 1 Author"
	Book1.subject = "Book 1 Subject"
	printbook(Book1)
}

func printbook(book Books) {
	fmt.Printf("%s\n", book.title)
	fmt.Printf("%s\n", book.author)
	fmt.Printf("%s\n", book.subject)
}
