package main

import "fmt"

// Defining a Structure
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	// Accessing Structure Members
	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	/* print Book1 info */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* print Book2 info */
	printBook(Book2)
	pointerBook(&Book2)
}

// Structures as Function Arguments

func printBook(book Books) {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}

// Pointers to Structures

func pointerBook(book *Books){
	fmt.Printf( "Book title : %s\n", book.title);
   	fmt.Printf( "Book author : %s\n", book.author);
   	fmt.Printf( "Book subject : %s\n", book.subject);
   	fmt.Printf( "Book book_id : %d\n", book.book_id);
}


