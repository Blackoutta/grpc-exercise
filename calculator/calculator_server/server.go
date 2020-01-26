package main

import (
	"context"
	"fmt"
	"io"
	calcpb "learnings/grpc/calculator/calculator_pb"
	"learnings/grpc/errors"
	"log"
	"net"
	"time"

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

func (s *server) Decompose(req *calcpb.DecomposeRequest, stream calcpb.CalculatorService_DecomposeServer) error {
	n := req.GetNumber()
	k := int64(2)
	for n > 1 {
		if n%k == 0 {
			result := &calcpb.DecomposeResponse{
				Result: k,
			}
			if err := stream.Send(result); err != nil {
				log.Fatalln("error while sending Decompose RPC call:", err)
			}
			n = n / k
			time.Sleep(time.Second)
			continue
		}
		k++
	}
	return nil
}

func (s *server) Average(stream calcpb.CalculatorService_AverageServer) error {
	fmt.Println("Average function was invoked! ")
	var sum int32
	var count int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := float32(sum) / float32(count)
			resp := &calcpb.AverageResponse{
				Result: result,
			}
			return stream.SendAndClose(resp)
		}
		errors.HandleError("error while receiving req from the client stream", err)
		sum += req.GetNumber()
		count++
	}
}

func main() {
	// port binding and GRPC server initialization
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// adding handlers
	calcpb.RegisterCalculatorServiceServer(s, &server{})

	// start server
	fmt.Println("greet server started, listening on 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
