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

	library.ListAllBooks()
	fmt.Println(library.FindBookById(0))
	fmt.Println(library.FindBookById(10))
	fmt.Println(library.FindBookByTitle("Golang"))
	fmt.Println(library.FindBookByTitle("ExampleTitle2"))

}
