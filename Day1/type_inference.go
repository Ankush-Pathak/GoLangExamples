package main

import "fmt"

func main2() {
	var address = "Pune" //This has become string now, initialization required if type not given
	fmt.Println(address)
	//address = 1 not possible
	fmt.Printf("%T\n", address) // Print the type

	var id, location, status = 1, "Pune", true //Multiple types initialized in one go

	fmt.Println(id, location, status)
}
