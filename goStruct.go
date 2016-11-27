package main

import (
	"errors"
	"fmt"
)

type AuthorName struct {
	first_name string
	last_name  string
}

type Book struct {
	AuthorName
	title   string
	book_id int
}

func CreateBook(first_name, last_name, book_title string, id int) (Book, error) {
	var book Book
	if (first_name == "") || (last_name == "") || (book_title == "") || (id < 0) {
		return Book{}, errors.New("Cannot create a new book.")
	} else {
		book = Book{AuthorName{first_name, last_name}, book_title, id}
		return book, errors.New("Cannot create a new book.")
	}
}

func ModifyBook(b *Book) {
	b.first_name = "J"
	b.last_name = "B"
	b.book_id = 99
	fmt.Printf("Inside modification: %T %p %v\n", b, &b, b)
}

// Method with a value receiver
func (b Book) MethodModifyBookByValue() {
	b.first_name = "X"
	b.last_name = "Y"
	b.book_id = 111
	fmt.Printf("Inside modification: %T %p %v\n", b, &b, b)
}

// Method with a pointer receiver
func (b *Book) MethodModifyBookByPointer() {
	b.first_name = "X"
	b.last_name = "Y"
	b.book_id = 111
	fmt.Printf("Inside modification: %T %p %v\n", b, &b, b)
}

func main() {
	fmt.Println("Hello Go!!! My good friend:)")
	book_1 := Book{}

	book_2 := Book{
		AuthorName{},
		"Book 2",
		5}
	book_3 := Book{
		AuthorName{
			"Domino",
			"Stefano"},
		"Golang Book",
		1000} // No comma, the bracket has to be in the sma line
	book_4 := Book{
		AuthorName: AuthorName{
			first_name: "D",
			last_name:  "S"},
		title:   "GB",
		book_id: 1, // Here is a comma, so the bracket can be in the next line
	}

	book_5, _ := CreateBook("Norek", "Fikolek", "Kotek na plotek", 7)
	book_6 := &Book{AuthorName{}, "Book 2", 5}

	book_7 := &Book{}
	book_8 := new(Book)

	fmt.Println(book_1)             // {{ }  0}
	fmt.Println(book_2)             // {{ } Book 2 5}
	fmt.Println(book_3)             // {{Domino Stefano} Golang Book 1000}
	fmt.Println(book_4)             // {{D S} GB 1}
	fmt.Println(book_5)             // {{Norek Fikolek} Kotek na plotek 7}
	fmt.Printf("Type %T\n", book_5) // Type main.Book
	fmt.Printf("Type %T\n", book_6) // Type *main.Book
	fmt.Printf("Type %T\n", book_7) // Type *main.Book
	fmt.Printf("Type %T\n", book_8) // Type *main.Book
	fmt.Println(*book_7 == *book_8) // True

	fmt.Printf("\nBefore modification:  %T %p  %v\n", book_1, &book_1, book_1)
	ModifyBook(&book_1)
	fmt.Printf("After modification:   %T %p  %v\n", book_1, &book_1, book_1)

	fmt.Printf("\nBefore modification: %T %p %v\n", book_7, &book_7, book_7)
	ModifyBook(book_7)
	fmt.Printf("After modification:  %T %p %v\n", book_7, &book_7, book_7)

	fmt.Printf("\nBefore modification:  %T %p %v\n", book_1, &book_1, book_1)
	book_7.MethodModifyBookByPointer()
	fmt.Printf("After modification:   %T %p %v\n", book_1, &book_1, book_1)

	fmt.Printf("\nBefore modification: %T %p %v\n", book_7, &book_7, book_7)
	book_7.MethodModifyBookByPointer()
	fmt.Printf("After modification:  %T %p %v\n", book_7, &book_7, book_7)
}
