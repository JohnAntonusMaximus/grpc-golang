package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	calculator "github.com/johnantonusmaximus/grpc-golang/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

func (s *server) SquareRoot(ctx context.Context, req *calculator.SquareRootRequest) (*calculator.SquareRootResponse, error) {
	number := req.GetFactor()
	fmt.Printf("Received Factor: %v\n", number)

	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received negative number: %v\n", number))
	}

	return &calculator.SquareRootResponse{Root: math.Sqrt(float64(number))}, nil

}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Server Error: %v", err)
	}

	s := grpc.NewServer()

	calculator.RegisterCalculatorServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Listener Error: %v", err)
	}

}
