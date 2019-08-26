package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/johnantonusmaximus/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Called: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	return &greetpb.GreetResponse{
		Result: "Hello " + firstName + "!",
	}, nil
}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
