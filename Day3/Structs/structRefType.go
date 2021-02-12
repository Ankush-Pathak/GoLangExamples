package main

import "fmt"

type Product1 struct {
	id   int
	name string
}

func main() {
	prod := new(Product1)
	prod.id = 10
	prod.name = "A"

	fmt.Println(prod)
	fmt.Println(*prod)

	changeProduct1(prod)

	//Alt way
	prod1 := &Product1{1, "A"}
	_ = prod1

	fmt.Println(prod)
}

func changeProduct1(product *Product1) {
	product.name = "B"
}
