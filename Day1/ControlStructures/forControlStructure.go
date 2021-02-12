package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	//While
	var j = 0
	for j < 10 {
		fmt.Println("j:", j)
		j++
	}

	var k = 0
	//Infinite loop
	for {
		if k > 3 {
			break
		}
		fmt.Println("Go!")
		k++
	}
}
