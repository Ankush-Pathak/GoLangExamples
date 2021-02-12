package main

import "fmt"

// Arrays are fixed len
// Slices are dynamic arrays

func main() {
	var array [3]int //Value type, so would be initialized to zero

	fmt.Println(array)

	array[1] = 100
	array[2] = 200

	for j := 0; j < len(array); j++ {
		fmt.Println(array[j])
	}

	for index, value := range array {
		fmt.Println("arra", index, ":", value)
	}

	var secondArray = [3]int{1, 2, 3}

	fmt.Println(secondArray)
}
