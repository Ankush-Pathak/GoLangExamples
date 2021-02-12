package main

import "fmt"

func main() {
	greet := func(id int) string {
		return "Hello"
	}

	fmt.Println(greet(100))

	process(12, postProcessing)
	process(20, func(id int) {
		fmt.Println("Please try again later, salary not processed yet.")
	})

}

func process(id int, postProcessing func(int)) {
	switch {
	case id > 10:
		postProcessing(10000)
	}
}

func postProcessing(sal int) {
	fmt.Println("Computed salary is", sal)
}
