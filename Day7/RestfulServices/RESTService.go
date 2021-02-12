package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"RESTfulServices/DAL"

	"github.com/gorilla/mux"
)

func getBooks(writer http.ResponseWriter, req *http.Request) {
	books := DAL.GetBooks()
	json.NewEncoder(writer).Encode(books)
}

func getBookByIsbn(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	isbn := params["isbn"]
	book := DAL.GetBookByIsbn(isbn)

	if book == (DAL.Book{}) {
		http.Error(writer, "Resource not available", 404)
		return
	}
	json.NewEncoder(writer).Encode(book)
}

func postBook(writer http.ResponseWriter, req *http.Request) {
	var book DAL.Book
	error := json.NewDecoder(req.Body).Decode(&book)
	if error != nil {
		fmt.Println("Error:", error)
	}
	fmt.Println("Book:", book)
	DAL.InsertBook(book)
	writer.WriteHeader(http.StatusCreated)
}

func deleteBookByIsbn(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	isbn := params["isbn"]
	DAL.DeleteBookByIsbn(isbn)
}

func sampleMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fmt.Println("Waring in the middle")
		next.ServeHTTP(writer, req)
		fmt.Println("Post handling middling")
	})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", sampleMiddleWare(getBooks)).Methods("GET")
	router.HandleFunc("/books/{isbn}", getBookByIsbn).Methods("GET")
	router.HandleFunc("/books/{isbn}", deleteBookByIsbn).Methods("DELETE")
	router.HandleFunc("/books", postBook).Methods("POST")
	http.ListenAndServe(":8000", router)
}
