package main

import "fmt"

func main() {
	type Employee struct {
		fName string
		lName string
		id    int
	}

	emp := Employee{fName: "A", lName: "B", id: 1} //Without keys, order matters
	fmt.Println(emp)
	fmt.Println(emp.fName)
	fmt.Println(emp.id)

	var emp1 Employee

	fmt.Println(emp1)
}
