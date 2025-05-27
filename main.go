package main

import (
	"fmt"
	"solnijko/golibrary/library"
)

func main() {
	library.AddBook(
		"ExampleTitle",
		"Example description for the book",
		false,
	)
	library.AddBook(
		"ExampleTitle2",
		"Example description for the book 2",
		true,
	)

	fmt.Println("List all books test")
	library.ListAllBooks()

	fmt.Println("Book search test")
	fmt.Println(library.FindBookById(0))
	fmt.Println(library.FindBookById(10))
	fmt.Println(library.FindBookByTitle("Golang"))
	fmt.Println(library.FindBookByTitle("ExampleTitle2"))

	fmt.Println("Book borrow and return test")
	borrowedBook, _ := library.BorrowBook("ExampleTitle2")
	fmt.Println(borrowedBook)
	library.ListAllBooks()
	library.ReturnBook(borrowedBook)
	library.ListAllBooks()
	fmt.Println(library.ReturnBook(nil))
}
