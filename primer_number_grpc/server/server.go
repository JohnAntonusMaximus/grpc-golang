package main

import (
	"fmt"
	"log"
	"net"

	prime "github.com/johnantonusmaximus/grpc-golang/primer_number_grpc/prime_number_pb"
	"google.golang.org/grpc"
)

func (s *server) GetPrimeFactors(req *prime.PrimeStreamRequest, stream prime.PrimeStreamService_GetPrimeFactorsServer) error {
	n := req.GetFactor()
	var k int32 = 2
	var f []int32

	for n > 1 {
		if n%k == 0 {
			log.Printf("Factor found: %v\n", k)
			resp := &prime.PrimeStreamResponse{
				Result: k,
			}
			f = append(f, k)
			stream.Send(resp)
		}
		if checkFactorsFound(f, n) == true {
			break
		}
		k = k + 1
		continue

	}
	return nil
}

type server struct{}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error creating listener: %v\n", err)
	}

	s := grpc.NewServer()

	prime.RegisterPrimeStreamServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to start stream server: %v\n", err)
	}
}

func checkFactorsFound(f []int32, n int32) bool {
	var result int32 = 1

	fmt.Println("F: ", f)

	for _, v := range f {
		result *= v
	}

	fmt.Println("Result: ", result)

	if result == n || result > n {
		return true
	}

	return false

}
