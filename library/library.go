package library

import "fmt"

type Book struct {
	Id                 int
	Title, Description string
	IsRead, IsBorrowed bool
}

var Books = make(map[int]*Book)
var BooksByTitle = make(map[string]*Book)
var NewBookId = 0

func AddBook(newTitle string, newDescription string, isRead bool) {
	newBook := &Book{
		Id:          NewBookId,
		Title:       newTitle,
		Description: newDescription,
		IsRead:      isRead,
		IsBorrowed:  false,
	}
	// Books = append(Books, newBook)
	BooksByTitle[newTitle] = newBook
	Books[NewBookId] = newBook
	NewBookId += 1
}

// Output book slice as a string
func (book *Book) String() string {
	return fmt.Sprintf("id: %d, title: %s, description: %s, isRead: %v, isBorrowed: %v",
		book.Id, book.Title, book.Description, book.IsRead, book.IsBorrowed)
}

func ListAllBooks() {
	for _, book := range Books {
		fmt.Println(book)
	}
}

func FindBookById(book_id int) (*Book, error) {
	if book, found := Books[book_id]; found {
		return book, nil
	}
	return nil, fmt.Errorf("book with such id was not found")
}

func FindBookByTitle(title string) (*Book, error) {
	if book, found := BooksByTitle[title]; found {
		return book, nil
	}
	return nil, fmt.Errorf("book with such title was not found")
}

func BorrowBook(book_title string) (*Book, error) {
	book, err := FindBookByTitle(book_title)
	if book != nil {
		book.IsBorrowed = true
	}
	return book, err
}

func ReturnBook(book *Book) error {
	if book == nil {
		return fmt.Errorf("cannot return nonexistence to library")
	}
	book.IsBorrowed = false
	return nil
}
