package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/johnantonusmaximus/grpc-golang/bidirectional_stream/bidirectionalpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection to GRPC failed: %v", err)
	}

	defer conn.Close()

	c := bidirectionalpb.NewGreetServiceClient(conn)

	doBidirectionalStream(c)
}

func doBidirectionalStream(c bidirectionalpb.GreetServiceClient) {
	fmt.Println("Starting to do a bidirectional streaming RPC...")
	ctx := context.Background()

	stream, err := c.GreetEveryone(ctx)
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v\n", err)
		return
	}

	ch := make(chan struct{})

	go func() {
		for _, req := range Requests {
			stream.Send(req)
			time.Sleep(time.Second * 1)
		}
		stream.CloseSend()

	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received From Server: %v\n", res.GetMessage())
		}
		close(ch)
	}()

	<-ch
}
