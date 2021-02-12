package main

import "fmt"

type Product struct {
	name  string
	value int
}

func main() {

	p := Product{"A", 10}

	changeProduct(&p)

	fmt.Println(p)
}

func changeProduct(p *Product) {
	p.value = 100
}
