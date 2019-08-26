package main

import (
	"context"
	"log"

	"github.com/johnantonusmaximus/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection to GRPC failed: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	ctx := context.Background()
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName:  "Radosta",
		},
	}

	resp, err := c.Greet(ctx, req)
	if err != nil {
		log.Fatalf("Greet RPC failed: %v ", err)
	}
	log.Printf("Response From Server: %v", resp)
}
