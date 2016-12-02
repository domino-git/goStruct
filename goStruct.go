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
	book_id string
}

func CreateBook(first_name, last_name, book_title, id string) (Book, error) {
	var book Book
	if (first_name == "") || (last_name == "") || (book_title == "") || (id == "") {
		return Book{}, errors.New("Cannot create a new book.")
	} else {
		book = Book{AuthorName{first_name, last_name}, book_title, id}
		return book, errors.New("Cannot create a new book.")
	}
}

func ModifyBook(b *Book) {
	b.first_name = "J"
	b.last_name = "B"
	b.book_id = "xxxx"
	//fmt.Printf("Inside modification: %T %p %v\n", b, &b, b)
}

// Method with a value receiver
func (b Book) MethodModifyBookByValue() {
	b.first_name = "X"
	b.last_name = "Y"
	b.book_id = "x"
	//fmt.Printf("Inside modification: %T %p %v\n", b, &b, b)
}

// Method with a pointer receiver
func (b *Book) MethodModifyBookByPointer() {
	b.first_name = "X"
	b.last_name = "Y"
	b.book_id = "x"
	//fmt.Printf("Inside modification: %T %p %v\n", b, &b, b)
}

func PrintDetails(object_id string, b Book) {
	fmt.Println(object_id, " ", b)
}

func main() {
	fmt.Println("Hello Go!!! My good friend:)")
	book_1 := Book{}

	book_2 := Book{
		AuthorName{},
		"Boooook2",
		"book_2"}
	book_3 := Book{
		AuthorName{
			"Domino",
			"Stefano"},
		"GolangBook",
		"book_3"} // No comma, the bracket has to be in the sma line
	book_4 := Book{
		AuthorName: AuthorName{
			first_name: "D",
			last_name:  "S"},
		title:   "GB",
		book_id: "book_4", // Here is a comma, so the bracket can be in the next line
	}

	book_5, _ := CreateBook("Norek", "Fikolek", "KotekNaPlotek", "book_5")
	book_6 := &Book{AuthorName{}, "Booook6", "book_6"}

	book_7 := &Book{}
	book_8 := new(Book)

	PrintDetails("book_1 ", book_1)  // {{"" ""} "" ""}
	PrintDetails("book_2 ", book_2)  // {{"" ""} "Boooook2" "book_2"}
	PrintDetails("book_3 ", book_3)  // {{"Domino" "Stefano"} "GolangBook" "book_3"}
	PrintDetails("book_4 ", book_4)  // {{"D" "S"} "GB" "book_4"}
	PrintDetails("book_5 ", book_5)  // {{"Norek" "Fikolek"} "KotekNaPlotek" "book_5"}
	PrintDetails("book_6 ", *book_6) // {{"" ""} "Booook6" "book_6}
	PrintDetails("book_7 ", *book_7) // {{"" ""} "" ""}
	PrintDetails("book_8 ", *book_8) // {{"" ""} "" ""}

	fmt.Printf("Type %T\n", book_5) // Type main.Book  - ordinary value
	fmt.Printf("Type %T\n", book_6) // Type *main.Book - pointer
	fmt.Printf("Type %T\n", book_7) // Type *main.Book - pointer
	fmt.Printf("Type %T\n", book_8) // Type *main.Book - pointer
	fmt.Println(*book_7 == *book_8) // True	- because pointer point to empty structs

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
