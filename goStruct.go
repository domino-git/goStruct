package main

import "fmt"

type AuthorName struct {
	first_name string
	last_name  string
}

type Book struct {
	AuthorName
	title   string
	book_id int
}

func main() {
	fmt.Println("Hello Go!!! My good friend:)")
	book_1 := Book{}
	fmt.Println(book_1)
	book_2 := Book{
		AuthorName{"Domino", "Stefano"},
		"Golang Book",
		1000}
	book_3 := Book{
		AuthorName: AuthorName{first_name: "D", last_name: "S"},
		title:      "GB",
		book_id:    1,
	}
	fmt.Println(book_2)
	fmt.Println(book_3)
}
