package main

import "fmt"

func main() {
	slice := []int{10, 20, 30}

	slice = append(slice, 40, 50)

	fmt.Println(slice)
}
