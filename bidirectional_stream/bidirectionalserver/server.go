package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/johnantonusmaximus/grpc-golang/bidirectional_stream/bidirectionalpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GreetEveryone(stream bidirectionalpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function was invoked by streaming request...\n")

	for {
		req, err := stream.Recv()
		firstName := req.GetGreeting().GetFirstName()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v\n", err)
			return err
		}
		log.Printf("Received from Client: %v\n", firstName)
		msg := "Hello there, " + firstName + "!"
		err = stream.Send(&bidirectionalpb.GreetEveryoneResponse{
			Message: msg,
		})
		if err != nil {
			log.Fatalf("Error sending response: %v\n", err)
			return err
		}
	}
}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()

	bidirectionalpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
