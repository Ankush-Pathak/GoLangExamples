package main

import (
	"fmt"
	"strconv"
)

//Or import on multiple lines would also work

func main() {
	var name = "name"
	var val = "12345"

	value, err := strconv.Atoi(name)
	fmt.Println(value, err)

	value, err = strconv.Atoi(val)
	fmt.Println(value, err)

}
