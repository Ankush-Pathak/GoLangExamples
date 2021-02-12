package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //For synchronization

func CheckoutProducts() {
	fmt.Println("Starting checkout...")
	time.Sleep(5 * time.Second)
	fmt.Println("Checkout done...")
	wg.Done()
}

func BrowseProducts() {
	fmt.Println("Starting browsing...")
	time.Sleep(5 * time.Second)
	fmt.Println("Browsing done...")
	wg.Done()
}

func main() {
	wg.Add(2)             //No of routines that will be run
	go CheckoutProducts() //Go routine
	go BrowseProducts()
	wg.Wait() //Will wait for two dones
	//time.Sleep(10 * time.Second) // Needed as program may exit before routines are scheduled and executed, if wg not used
}
