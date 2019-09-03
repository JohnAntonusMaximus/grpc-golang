package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/johnantonusmaximus/grpc-golang/long_greet_stream/greetstreampb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection to GRPC failed: %v", err)
	}

	defer conn.Close()

	c := greetstreampb.NewGreetServiceClient(conn)

	doClientStream(c)
}

func doClientStream(c greetstreampb.GreetServiceClient) {
	fmt.Println("Starting to do a client streaming RPC...")
	ctx := context.Background()

	requests := []*greetstreampb.LongGreetRequest{
		&greetstreampb.LongGreetRequest{
			Greeting: &greetstreampb.Greeting{
				FirstName: "Stephanie",
			},
		},
		&greetstreampb.LongGreetRequest{
			Greeting: &greetstreampb.Greeting{
				FirstName: "Mark",
			},
		},
		&greetstreampb.LongGreetRequest{
			Greeting: &greetstreampb.Greeting{
				FirstName: "Peter",
			},
		},
		&greetstreampb.LongGreetRequest{
			Greeting: &greetstreampb.Greeting{
				FirstName: "Harold",
			},
		},
		&greetstreampb.LongGreetRequest{
			Greeting: &greetstreampb.Greeting{
				FirstName: "John",
			},
		},
		&greetstreampb.LongGreetRequest{
			Greeting: &greetstreampb.Greeting{
				FirstName: "Chris",
			},
		},
	}

	stream, err := c.LongGreetStream(ctx)
	if err != nil {
		log.Fatalf("Error while calling long greet: %v\n", err)
	}

	for _, req := range requests {
		stream.Send(req)
		time.Sleep(time.Second * 1)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}
	log.Printf("Response: %v\n", resp)
}
