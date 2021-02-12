package main

import "fmt"

func main() {
	//Interpreted string
	var msg = "Welcome to GoLang training"

	//Raw string
	var title = `Day 1`
	var description = `You just learned that this char allows
    multiple lines but not \n interpretation`
	var multiIntr = "You can have multiple \nlines as so here also"
	fmt.Println(msg)
	fmt.Println(title)
	fmt.Println(description)
	fmt.Println(multiIntr)
}
