package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"learnings/grpc/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetingResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	// port binding and GRPC server initialization
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// adding handlers
	greetpb.RegisterGreetServiceServer(s, &server{})

	// start server
	fmt.Println("greet server started, listening on 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
