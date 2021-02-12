package main

import (
	"fmt"
	"sync"
)

var val = 0

func worker(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	val++
	mutex.Unlock()

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &mutex)
	}

	wg.Wait()
	fmt.Println("VAL:", val) //Will always be 1000

}
