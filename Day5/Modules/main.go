package main

//Module created using "go mod init myproject"

import (
	"myproject/DataAccess"
	DAM "myproject/DataAccess/Messaging" //Using aliases to resolve ambiguity
	XM "myproject/XML/Messaging"
)

func main() {
	DataAccess.ReadDB()
	DAM.SendMsg()
	XM.RecvMsg()
}
