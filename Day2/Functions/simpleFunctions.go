package main

import "fmt"

func main() {
	func1(1, "John")
}

func func1(id int, name string) {
	fmt.Println("func1 is called...", id, name)
}

func func2(id, sal int) { //Same type specification can be grouped
	fmt.Println(id, sal)
}

func func3(id, sal int) int { //Returning int
	return id + sal
}

func func4(id, sal int) (int, int) { //Returning multiple values
	return id, sal
}

func func5(id, sal int) (bonus int, status bool) { // Named return, these variables are declared due to this and no need to initialize them
	bonus = sal + 1000
	status = true
	return
}
