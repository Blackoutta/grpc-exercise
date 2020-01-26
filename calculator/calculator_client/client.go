package main

import (
	"context"
	"fmt"
	calcpb "learnings/grpc/calculator/calculator_pb"
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
	c := calcpb.NewSumServiceClient(cc)

	// invoking grpc calls
	doUnary(c)
}

func doUnary(c calcpb.SumServiceClient) {
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
