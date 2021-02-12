package main

import "fmt"

type User struct {
	name string
}

type Admin struct {
	User
	Message
	role int
}

type Message struct {
	msg string
}

func (msg Message) login() {
	fmt.Println("Logged in into msg platform")
}

func (user User) login() {
	fmt.Println("User logged in")
}

func (admin Admin) login() {
	fmt.Println("Admin logged in")
}

func main() {
	user := User{"A"}
	msg := Message{"B"}
	admin := Admin{user, msg, 1}

	//admin.login() is ambiguous if Admin::login is not defined
	admin.login()
	admin.User.login()
	admin.Message.login()

	fmt.Println(admin)
}
