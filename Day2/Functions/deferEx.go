package main

import "fmt"

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func main() {
	defer foo()

	bar()
}
