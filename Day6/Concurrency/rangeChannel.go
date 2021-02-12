package main

import "fmt"

func Producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //So that range can loop over it
}

func main() {
	ch := make(chan int)
	go Producer(ch)

	for val := range ch { //Range for channel returns only val, not index
		fmt.Println("Val", val)
	}
}
