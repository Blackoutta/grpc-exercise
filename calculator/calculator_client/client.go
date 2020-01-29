package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	calcpb "learnings/grpc/calculator/calculator_pb"
	"learnings/grpc/errors"
	"log"
	"sync"
	"time"

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
	//doClientStream(c)
	//doBiDirectionalStream(c)
	doErrorUnary(c)
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

func doBiDirectionalStream(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Bi-directional Stream RPC!")

	// get stream
	stream, err := c.FindMaximum(context.Background())
	errors.HandleError("error while calling FindMaximum RPC", err)

	requests := []*calcpb.FindMaximumRequest{
		{
			Number: 1,
		},
		{
			Number: 6,
		},
		{
			Number: 2,
		},
		{
			Number: 8,
		},
		{
			Number: 5,
		},
		{
			Number: 66,
		},

	}
	var wg sync.WaitGroup

	wg.Add(1)
	// send requests
	go func() {
		for _, req := range requests {
			err := stream.Send(req)
			errors.HandleError("error while sending requests to server", err)
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		errors.HandleError("error while closing the sending to server", err)
	}()

	// receive response
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				return
			}
			errors.HandleError("error while receiving from server", err)
			fmt.Printf("FindMaximum Response: %v\n", resp.GetResult())
		}
	}()


	wg.Wait()
}

func doErrorUnary(c calcpb.CalculatorServiceClient) {
	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -10)

}

func doErrorCall(c calcpb.CalculatorServiceClient, n int32) {
	resp, err := c.SquareRoot(context.Background(), &calcpb.SquareRootRequest{
		Number: n,
	})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from GRPC (user error)
			fmt.Printf("error message from server: %v", respErr.Message())
			fmt.Printf("error code from server: %v\n", respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("client: we probably sent a negative number.")
				return
			}
		} else {
			log.Fatalf("Big error calling SquareRoot: %v\n", err)
		}
	}
	fmt.Printf("Result of square root of %v: %v\n", n, resp.GetNumberRoot())
}