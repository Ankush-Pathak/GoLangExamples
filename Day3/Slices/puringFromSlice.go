package main

import "fmt"

func main() {
	slice := []string{"MON", "TUE", "WED", "THU"}

	slice = append(slice[1:2], slice[2:]...) //Remove first element

	fmt.Println(slice)
}
