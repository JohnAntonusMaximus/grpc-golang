package main

import (
	"context"
	"io"
	"log"

	prime "github.com/johnantonusmaximus/grpc-golang/primer_number_grpc/prime_number_pb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect to server: %v\n", err)
	}

	defer conn.Close()

	c := prime.NewPrimeStreamServiceClient(conn)

	doStream(c)
}

func doStream(c prime.PrimeStreamServiceClient) {

	var f int32 = 160

	req := &prime.PrimeStreamRequest{
		Factor: f,
	}

	ctx := context.Background()

	log.Printf("Factoring %v...", f)

	resStream, err := c.GetPrimeFactors(ctx, req)
	if err != nil {
		log.Fatalf("Error response from server: %v\n", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v\n", err)
		}
		log.Printf("Received Factor: %v\n", msg)
	}

	log.Printf("All factors received for %v!\n", f)

}
