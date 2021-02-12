package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"my_microservice/greetpb"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type ServerType struct {
}

func (*ServerType) Greet(context context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fname := request.GetGreeting().GetFirstName()
	lname := request.GetGreeting().GetLastName()

	result := fname + " " + lname

	response := &greetpb.GreetResponse{Result: result}

	return response, nil
}

func (*ServerType) CalcSqRoot(context context.Context, request *greetpb.SquareRootRequest) (*greetpb.SquareRootResponse, error) {
	// Retrive data from request
	number := request.GetNumber()
	sqroot := math.Sqrt(float64(number))
	response := &greetpb.SquareRootResponse{NumberRoot: sqroot}
	return response, nil
}

func (*ServerType) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("Greeting many times for", req)
	firstName := req.GetGreeting().FirstName
	lastName := req.GetGreeting().LastName

	for i := 1; i <= 10; i++ {
		result := "Hello " + firstName + " " + lastName + " no: " + strconv.Itoa(i)
		resp := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(resp)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	listener, error := net.Listen("tcp", "0.0.0.0:8000")
	if error != nil {
		log.Fatal(error)
	}

	server := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(server, &ServerType{})

	err := server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
