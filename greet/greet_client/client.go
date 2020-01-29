package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"learnings/grpc/errors"
	"learnings/grpc/greet/greetpb"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// boilerplate
	fmt.Println("hello I'm a grpc client!")

	certFile := "ssl/ca.crt" // certificate authority trust certificate
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	errors.HandleError("error while loading ca trust certificate", sslErr)

	opts := grpc.WithTransportCredentials(creds)
	cc, err := grpc.Dial("localhost:50051", opts)
	defer cc.Close()
	if err != nil {
		log.Fatalln("could not connect:", err)
	}
	c := greetpb.NewGreetServiceClient(cc)

	// invoking grpc calls
	doUnary(c)
	// doServerStream(c)
	// doClientStream(c)
	//doBiDirectionalStream(c)
	//doUnaryWithDeadline(c, 5) // should complete
	//doUnaryWithDeadline(c, 1) // should timeout

}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Starting to do Unary RPC with deadline!")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "yang",
			LastName:  "hu",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout * time.Second)
	defer cancel()
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded!")
				fmt.Printf("error message: %v\n", statusErr.Message())
			} else {
				fmt.Printf("unexpected error: %v", statusErr)
			}
			return
		} else {
			log.Fatalf("fatal error while calling Greet RPC: %v", statusErr)
		}
	}
	log.Printf("Response from Greet: %v", res.GetResult())
}


func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Unary RPC!")
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "yang",
			LastName:  "hu",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Printf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.GetResult())
}

func doServerStream(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Server Streaming RPC!")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "yang",
			LastName:  "hu",
		},
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling GreetManyTimes RPC:", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("error while reading stream:", err)
		}
		result := msg.GetResult()
		log.Printf("Response from GreetManyTimes: %v\n", result)
	}

}

func doClientStream(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Client Streaming RPC!")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Yang",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Rin",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lewis",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ben",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Shane",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	errors.HandleError("error while calling Long Greet", err)
	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req)
		err := stream.Send(req)
		errors.HandleError("error while sending LongGreet request", err)
		time.Sleep(100 * time.Microsecond)
	}
	res, err := stream.CloseAndRecv()
	errors.HandleError("error while receiving response from LongGreet", err)
	fmt.Printf("LongGreet Response: %v\n", res.GetResult())
}

func doBiDirectionalStream(c greetpb.GreetServiceClient) {
	var wg sync.WaitGroup
	fmt.Println("Starting to do Bi-directional Streaming RPC!")

	// create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	errors.HandleError("error while calling GreetEveryone RPC", err)

	wg.Add(1)

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Yang",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Rin",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lewis",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ben",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Shane",
			},
		},
	}

	// send a bunch of messages to the server
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			err := stream.Send(req)
			errors.HandleError("error while sending message to server", err)
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		errors.HandleError("error while closing sending to server", err)
	}()
	// receive a bunch of messages from the server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				return
			}
			errors.HandleError("error while receiving response from server", err)
			fmt.Printf("Received: %v\n", res.GetResult())
		}
	}()
	// block until everything is done
	wg.Wait()
}
