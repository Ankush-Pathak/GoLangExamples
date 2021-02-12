package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	getDetails(slice...)
}

//Variadic params can also take a slice
func getDetails(marks ...int) {
	fmt.Println(marks)
}
