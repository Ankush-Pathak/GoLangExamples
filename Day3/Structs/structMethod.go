package main

import "fmt"

type Person struct {
	fname string
	lname string
}

//per Person is a receiver type
func (per Person) getPersonDetails() {
	fmt.Println("Person Details")
}

func main() {
	per := Person{"A", "B"}
	per.getPersonDetails()
}
