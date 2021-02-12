package main

import (
	"io"
	"net/http"
)

func respond(responseWriter http.ResponseWriter, req *http.Request) {
	io.WriteString(responseWriter, "Hello, world.")
}

func main() {
	http.HandleFunc("/test", respond)

	http.ListenAndServe(":8000", nil)
}
