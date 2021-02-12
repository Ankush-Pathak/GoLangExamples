package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"my_microservice/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is client")
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	client := greetpb.NewGreetServiceClient(conn)

	doUnaryRequest(client)
	doServerStreaming(client)

}

func doServerStreaming(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ankush",
			LastName:  "Pathak",
		},
	}

	stream, err := client.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err2 := stream.Recv()

		if err2 == io.EOF {
			fmt.Println("Done recv data from server.")
			return
		} else if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("Res from stream:", msg.GetResult())
	}
}

func doUnaryRequest(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ankush",
			LastName:  "Pathak",
		},
	}

	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GreetResponse:", res.Result)

	req2 := &greetpb.SquareRootRequest{
		Number: 16,
	}

	res2, err2 := client.CalcSqRoot(context.Background(), req2)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println("CalcSqRootResponse:", res2.NumberRoot)

}
