package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/johnantonusmaximus/grpc-golang/greet_stream_server/greetstreampb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection to GRPC failed: %v", err)
	}

	defer conn.Close()

	c := greetstreampb.NewGreetServiceClient(conn)

	doStream(c)
}

func doStream(c greetstreampb.GreetServiceClient) {

	req := &greetstreampb.GreetManyTimesRequest{
		Greeting: &greetstreampb.Greeting{
			FirstName: "John",
			LastName:  "Radosta",
		},
	}

	ctx := context.Background()

	resStream, err := c.GreetManyTimes(ctx, req)
	if err != nil {
		log.Fatalf("Greet RPC failed: %v ", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v ", err)
		}
		fmt.Printf("Stream: %v\n", msg.GetResult())
	}

	fmt.Println("Finished reading stream!")

}
