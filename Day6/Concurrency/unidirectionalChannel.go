package main

import "fmt"

func main() {
	ch := make(chan string)
	go Routine(ch)

	ch <- "Data"
}

//Read-only channel
func Routine(ch <-chan string) {
	fmt.Println("Read", <-ch)
}
