package main

import "fmt"

func main() {
	var name = "John"

	//Primitve data types supported
	switch name {
	case "John":
		fmt.Println("John")
		fallthrough //Will execute next case also
	case "Doe":
		fmt.Println("Doe")
	case "Johnny", "Done": //Will match Johnny or Done
		fmt.Println("John Doe")
	default:
		fmt.Println("Default")
	}

	switch {
	case len(name) > 5:
		fmt.Println("Len greater than 5")
	}
}
