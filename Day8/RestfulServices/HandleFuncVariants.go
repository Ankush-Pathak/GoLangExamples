package main

import (
	"io"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

type Person struct {
	name string
	age  int
}

func (per *Person) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "Hello, world.")
}

func main() {
	per := &Person{"A", 1}
	router := mux.NewRouter()

	router.Handle("/person", per)
	http.ListenAndServe(":8000", router)
	fmt.Println("Done")

}
