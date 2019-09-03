package main

import (
	"context"
	"fmt"
	"log"

	"github.com/johnantonusmaximus/grpc-golang/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error creating client connection: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	//doUnary(c)
	doErrorUnary(c)
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

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	ctx := context.Background()

	var factorOne int32 = 4
	var factorTwo int32 = -10

	req := &calculatorpb.SquareRootRequest{
		Factor: factorOne,
	}

	reqError := &calculatorpb.SquareRootRequest{
		Factor: factorTwo,
	}

	result, err := c.SquareRoot(ctx, req)
	if err != nil {
		log.Fatalf("Erorr with GRPC Client Invocation: %v\n", err)
	}

	log.Printf("Received From Server: %v\n", result.GetRoot())

	result, err = c.SquareRoot(ctx, reqError)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error gRPC (Business Logic error)
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number.")
				return
			}
		} else {
			log.Fatalf("Big error calling SquareRoot: %v\n", respErr)
			return
		}
	}

	log.Printf("Received From Server: %v\n", result)
}
