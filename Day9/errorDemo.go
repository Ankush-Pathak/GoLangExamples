package main

import "fmt"

func main() {
	OpenFile()
	fmt.Println("File opened...")
}

func OpenFile() {
	defer cleanUp()
	fmt.Println("Opening file...")

	panic("File not found")

	fmt.Println("Closing file...")
}

func cleanUp() {
	fmt.Println("Cleaning up...")
	recover() //Catches exception
}
