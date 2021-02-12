package main

import "fmt"

type Student struct {
	id   int
	name string
}

//Ref type recv
func (student *Student) getStudentDetails() {
	student.name = "B"
}

func main() {
	stu := Student{1, "A"}

	stu.getStudentDetails()

	fmt.Println(stu)
}
