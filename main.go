package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//structs (model)

//Book struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as a slice Book struct
var books []Book

//functions
func getBooks(w http.ResponseWriter, r *http.Request) {

}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init router
	r := mux.NewRouter()

	//Mock data - implement DB
	books = append(books, Book{ID: "1", Isbn: "44783", Title: "Book 1", Author: &Author{Firstname: "Ryan", Lastname: "Zola"}})
	books = append(books, Book{ID: "2", Isbn: "45532", Title: "Book 2", Author: &Author{Firstname: "J.D.", Lastname: "Salinger"}})
	books = append(books, Book{ID: "3", Isbn: "48291", Title: "Book 3", Author: &Author{Firstname: "Author", Lastname: "McBookwriter"}})
	books = append(books, Book{ID: "3", Isbn: "48291", Title: "Book 4", Author: &Author{Firstname: "Kaijun", Lastname: "MacInauthor"}})

	//Route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
