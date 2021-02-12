package main

import "fmt"

var getHits = func() func() int {
	var counter = 0

	var addHits = func() int {
		counter++
		return counter
	}

	return addHits
}

var getHitsSelfExec = func() func() int {
	var counter = 0

	var addHits = func() int {
		counter++
		return counter
	}

	return addHits
}() //Executes outer func, so getHitsSelfExec will contain the return value of outer function i.e addHits

func main() {
	fmt.Println(getHits()()) //1
	fmt.Println(getHits()()) //1

	fmt.Println(getHitsSelfExec()) //1
	fmt.Println(getHitsSelfExec()) //2, getHitsSelfExec is a closure, and it has access to counter
}
