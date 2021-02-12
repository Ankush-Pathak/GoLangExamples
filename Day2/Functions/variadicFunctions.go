package main

import "fmt"

//Explicit function overloading not available in go
func main() {
	calcMarks("John")
	calcMarks("ABC", 1)
	calcMarks("DEF", 1, 2)
}

func calcMarks(name string, marksSlice ...int) { //marks can be empty or have any no of elements, marks is a slice
	//Variadic params can only be the last param
	fmt.Println("Marks for", name, "are", marksSlice)
}

//func calcMarks(id string, name string) { // Will not work, overloading not allowed

//}
