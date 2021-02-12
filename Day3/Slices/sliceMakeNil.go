package main

import "fmt"

func main() {
	slice := make([]int, 4)

	fmt.Println(slice)

	slice = append(slice, 100, 300)

	fmt.Println(slice)

	fmt.Println("Len:", len(slice), "Cap:", cap(slice))

	//nil slice
	var slice2 []int

	_ = slice2
}
