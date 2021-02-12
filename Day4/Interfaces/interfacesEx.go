package main

import "fmt"

type IShape interface {
	area()
	perimeter()
}

func (triangle Triangle) area() {
	fmt.Println("Area:", triangle.side*triangle.side) //The formula is wrong, I know
}

func (triangle Triangle) perimeter() {
	fmt.Println("Peri:", 3*triangle.side)
}

type Triangle struct {
	side int
}

type Square struct {
	side int
}

func (square Square) area() {
	fmt.Println("SQ Area:", square.side*square.side)
}

func (square Square) perimeter() {
	fmt.Println("SQ Peri:", 4*square.side)
}

func main() {
	var shape IShape

	shape = Triangle{10} //Go implictly checks during this type cast if struct properly implements the interface
	shape.area()
	shape.perimeter()

	shape = Square{5}
	shape.area()
	shape.perimeter()
}
