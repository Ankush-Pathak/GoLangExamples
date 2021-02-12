package main

import "fmt"

func main() {
	ch := make(chan string, 2) //Buffered channel, can put two values without blocking
	ch <- "Go, lang!"
	ch <- "Come, lang!"
	//ch <- "This will block and kill the prog as all coroutines will be asleep"
	fmt.Println("Read", <-ch)
	ch <- "Stay, lang!"
	fmt.Println("Read", <-ch)

}
