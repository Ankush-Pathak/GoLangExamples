package main

import "fmt"

func main() {
	const pi = 3.14
	const (
		MON = iota //Auto-increment values, starts from 0
		TUE
		WED
		THURS = 6        //Also possible
		SAT   = iota + 1 //Break series and start from alternate value
		SUN
	)

	const (
		JAN = iota + 1 //Start from 1
	)

	fmt.Println(pi, MON, TUE, WED, THURS, SAT, SUN, JAN)
}
