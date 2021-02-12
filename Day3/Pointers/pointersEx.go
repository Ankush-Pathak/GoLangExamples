package main

import "fmt"

func main() {
	var i = 100

	ptr := &i

	changeVal(ptr)

	fmt.Println(i)

	array := [3]int{1, 2, 3}

	changeArray(&array)

	fmt.Println(array)

	slice := []int{4, 5, 6}

	changeSlice(slice) //Passed by ref only

	fmt.Println(slice)
}

func changeSlice(slice []int) {
	slice[2] = 600
}

func changeArray(array *[3]int) {
	array[2] = 300
}

func changeVal(ptr *int) {
	*ptr += 10
}
