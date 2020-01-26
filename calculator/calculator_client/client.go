package main

import (
	"context"
	"fmt"
	"io"
	calcpb "learnings/grpc/calculator/calculator_pb"
	"learnings/grpc/errors"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// boilerplate
	fmt.Println("hello I'm a grpc client!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer cc.Close()
	if err != nil {
		log.Fatalln("could not connect:", err)
	}
	c := calcpb.NewCalculatorServiceClient(cc)

	// invoking grpc calls
	// doUnary(c)
	// doServerStream(c)
	doClientStream(c)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary RPC!")
	req := &calcpb.SumRequest{
		Sum: &calcpb.Sum{
			Num1: 3,
			Num2: 10,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("err invoking Sum RPC")
	}
	fmt.Println("The result is: ", res.Result)
}

func doServerStream(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Server Stream RPC!")
	req := &calcpb.DecomposeRequest{
		Number: 65535,
	}
	stream, err := c.Decompose(context.Background(), req)
	if err != nil {
		log.Fatalln("error while sending Decompose RPC call:", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("error while reading stream:", err)
		}
		result := res.GetResult()
		fmt.Printf("Response from Decompose RPC: %v\n", result)
	}
}

func doClientStream(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Client Stream RPC!")
	requests := []*calcpb.AverageRequest{
		&calcpb.AverageRequest{
			Number: 1,
		},
		&calcpb.AverageRequest{
			Number: 2,
		},
		&calcpb.AverageRequest{
			Number: 3,
		},
		&calcpb.AverageRequest{
			Number: 4,
		},
	}
	stream, err := c.Average(context.Background())
	errors.HandleError("error while calling Average RPC", err)
	for _, req := range requests {
		fmt.Printf("Sending Average request: %v\n", req)
		err := stream.Send(req)
		errors.HandleError("error while sending Average request stream", err)
	}
	resp, err := stream.CloseAndRecv()
	errors.HandleError("error while receiving Average response", err)
	fmt.Printf("Average Response: %v\n", resp.GetResult())
}
