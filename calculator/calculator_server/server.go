package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calculator "github.com/johnantonusmaximus/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *calculator.CalculatorRequest) (*calculator.CalculatorResponse, error) {
	fmt.Printf("Sum Function Invoked: %v\n", req)
	integerOne := req.GetIntegerOne()
	integerTwo := req.GetIntegerTwo()
	result := integerOne + integerTwo

	return &calculator.CalculatorResponse{
		Result: result,
	}, nil
}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Server Error: %v", err)
	}

	s := grpc.NewServer()

	calculator.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Listener Error: %v", err)
	}

}
