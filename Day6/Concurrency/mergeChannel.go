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

func Merge(chSlice ...chan int) {
	for _, ch := range chSlice {
		fmt.Println("Val:", <-ch)
	}

}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go Server1(ch)
	go Server2(ch2)

	Merge(ch, ch2)
}
