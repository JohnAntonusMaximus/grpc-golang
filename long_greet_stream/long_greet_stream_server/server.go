package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/johnantonusmaximus/grpc-golang/long_greet_stream/greetstreampb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) LongGreetStream(stream greetstreampb.GreetService_LongGreetStreamServer) error {
	fmt.Printf("LongGreet function was invoked by streaming request...\n")

	for {
		req, err := stream.Recv()
		firstName := req.GetGreeting().GetFirstName()
		if err == io.EOF {
			msg := "Hello " + firstName + "! "
			return stream.SendAndClose(&greetstreampb.LongGreetResponse{
				Message: msg,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v\n", err)
		}
		log.Printf("Received from Client: %v", firstName)
	}
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
