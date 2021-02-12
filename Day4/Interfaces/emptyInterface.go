package main

import (
	"fmt"
	"reflect"
)

type Empty struct{}

func findType(i interface{}) { //Empty interace can accept any type
	//Println takes variadic empty interface as param
	fmt.Println("Taking param as any type")
	fmt.Println("Got type", reflect.TypeOf(i), "\n")
}

func main() {
	findType(100)
	findType("123")
	findType(Empty{})
}
