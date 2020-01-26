package main

import (
	"context"
	"fmt"
	calcpb "learnings/grpc/calculator/calculator_pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	num1 := req.GetSum().GetNum1()
	num2 := req.GetSum().GetNum2()
	sum := num1 + num2
	result := &calcpb.SumResponse{
		Result: sum,
	}
	return result, nil
}

func main() {
	// port binding and GRPC server initialization
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// adding handlers
	calcpb.RegisterSumServiceServer(s, &server{})

	// start server
	fmt.Println("greet server started, listening on 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
