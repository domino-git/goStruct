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

	fmt.Println(book_1)
	fmt.Println(book_2)
	fmt.Println(book_3)
	fmt.Println(book_4)
	fmt.Println(book_5)
}
