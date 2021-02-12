package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //For synchronization

func ReadDB(ch chan int) {
	fmt.Println("Starting reading...")
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Putting into channel", i)
		ch <- i //Will be blocked until this is read from other end, this is unbuffered
	}
	fmt.Println("Checkout done...")
	wg.Done()
}

func WriteFile(ch chan int) {
	fmt.Println("Starting browsing...")
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Data from channel", <-ch) //Will be blocked until there is something to read
	}
	fmt.Println("Browsing done...")
	wg.Done()
}

//Channels are ref types
func main() {
	ch := make(chan int)
	wg.Add(2)     //No of routines that will be run
	go ReadDB(ch) //Go routine
	go WriteFile(ch)
	wg.Wait() //Will wait for two wg.Done calls
	//time.Sleep(10 * time.Second) // Needed as program may exit before routines are scheduled and executed, if wg not used
}
