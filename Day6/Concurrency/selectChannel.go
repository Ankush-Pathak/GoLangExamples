package main

import (
	"fmt"
	"time"
)

func Server1(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 200
}

func Server2(ch chan int) {
	time.Sleep(4 * time.Second)
	ch <- 200
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go Server1(ch)
	go Server2(ch2)

	select { //Similar to switch case, used in context of go routines
	case res := <-ch:
		fmt.Println("Server1:", res)
	case res := <-ch2:
		fmt.Println("Server2:", res)
	case <-time.After(5 * time.Second):
		fmt.Println("Request timedout, no servers available")
	}
}
