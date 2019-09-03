package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/johnantonusmaximus/grpc-golang/greet_stream_server/greetstreampb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetstreampb.GreetRequest) (*greetstreampb.GreetResponse, error) {
	fmt.Printf("Greet Called: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	return &greetstreampb.GreetResponse{
		Result: "Hello " + firstName + "!",
	}, nil
}

func (s *server) GreetManyTimes(req *greetstreampb.GreetManyTimesRequest, stream greetstreampb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10000; i++ {
		result := "Hello " + firstName + "! #" + strconv.Itoa(i)
		res := &greetstreampb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(time.Millisecond * 10)
	}
	return nil
}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()

	greetstreampb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
