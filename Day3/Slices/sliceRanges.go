package main

import "fmt"

func main() {
	slice := []string{"MON", "TUE", "WED", "THU"}

	fmt.Println(slice[1:3]) //Upper value exclusive

	fmt.Println(slice[:2])

	fmt.Println(slice[2:])
}
