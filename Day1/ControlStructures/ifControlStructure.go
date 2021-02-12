package main

import "fmt"

//Only has if, else, for and switch

func main() {
	var i = 1000

	if i > 900 {
		fmt.Println(i, "is greater than 900")
	} else {
		fmt.Println("else")
	}

	if result := getDetails(); result < 1000 {
		//result available only in if
		fmt.Println(result, "is less than 1000")
	}
}

func getDetails() int { return 100 }
