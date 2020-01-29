package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	calcpb "learnings/grpc/calculator/calculator_pb"
	"learnings/grpc/errors"
	"log"
	"math"
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

func (s *server) FindMaximum(stream calcpb.CalculatorService_FindMaximumServer) error {
	fmt.Println("FindMaximum function was invoked! ")
	var nr []int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		errors.HandleError("error while receiving from client", err)
		nr = append(nr, req.GetNumber())
		sendErr := stream.Send(&calcpb.FindMaximumResponse{
			Result: findMax(nr),
		})
		errors.HandleError("error while sending response to client", sendErr)
	}
}

func (s *server) SquareRoot(ctx context.Context, req *calcpb.SquareRootRequest) (*calcpb.SquareRootResponse, error) {
	fmt.Println("SquareRoot function was invoked! ")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v\n", number),
		)
	}
	return &calcpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func findMax(n []int32) int32 {
	var max int32
	for i, v := range n {
		if i == 0 {
			max = v
			continue
		}
		if v > max {
			max = v
		}
	}
	return max
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

	// register reflection service on gRPC server
	reflection.Register(s)

	// start server
	fmt.Println("greet server started, listening on 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
