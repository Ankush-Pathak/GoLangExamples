package main

import "fmt"

func getComparator() func(a, b int) bool {
	comparator := func(a, b int) bool {
		if a > b {
			return true
		}
		return false
	}
	return comparator
}

func main() {
	comparator := getComparator()
	fmt.Println(comparator(10, 15))
	fmt.Println(comparator(20, 15))

}
