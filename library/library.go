package library

import "fmt"

type Book struct {
	id                 int
	title, description string
	isRead             bool
}

var Books []*Book
var BookId = 0

func AddBook(newTitle string, newDescription string, isRead bool) {
	newBook := &Book{
		id:          BookId,
		title:       newTitle,
		description: newDescription,
		isRead:      isRead,
	}
	Books = append(Books, newBook)
	BookId += 1
}

// Output book slice as a string
func (book *Book) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, Description: %s, IsRead: %v",
		book.id, book.title, book.description, book.isRead)
}

func ListAllBooks() {
	for _, book := range Books {
		fmt.Println(book)
	}
}

func FindBookById(book_id int) (*Book, error) {
	var booksByID = make(map[int]*Book)
	for _, book := range Books {
		booksByID[book.id] = book
	}
	if book, exists := booksByID[book_id]; exists {
		return book, nil
	}
	return nil, fmt.Errorf("Book with such id was not found")
}

func FindBookByTitle(book_title string) (*Book, error) {
	var booksByTitle = make(map[string]*Book)
	for _, book := range Books {
		booksByTitle[book.title] = book
	}
	if book, exists := booksByTitle[book_title]; exists {
		return book, nil
	}
	return nil, fmt.Errorf("Book with such title was not found")
}
