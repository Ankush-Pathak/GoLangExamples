package main

import (
	"fmt"
	"sync"
)

var val = 0

func worker(wg *sync.WaitGroup) {
	val += 1

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	fmt.Println("VAL:", val) //Wont be 1000 always

}
