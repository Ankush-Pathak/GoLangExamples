package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go LongRunningProcess(ch)
	fmt.Println("Control returned to main")
	val := <-ch
	DoSomething(val)

}

func DoSomething(val int) {
	fmt.Println(val, "Doing something...")
}

func LongRunningProcess(ch chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Long running task done!")
	ch <- 10
}
