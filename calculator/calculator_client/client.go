package main

import (
	"context"
	"log"

	"github.com/johnantonusmaximus/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error creating client connection: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	ctx := context.Background()

	var integerOne int32 = 3
	var integerTwo int32 = 10

	req := &calculatorpb.CalculatorRequest{
		IntegerOne: integerOne,
		IntegerTwo: integerTwo,
	}

	result, err := c.Sum(ctx, req)
	if err != nil {
		log.Fatalf("Erorr with GRPC Client Invocation: %v", err)
	}

	log.Printf("Summed Integers From Server: %v", result)
}
