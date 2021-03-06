package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"learnings/grpc/errors"
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

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("GreetManyTimes function was invoked!")
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		err := stream.Send(res)
		errors.HandleError("error while sending GreetManyTimes request", err)
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet function was invoked!")
	result := "Hello "

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// finished reading the stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("error while reading the client stream: %v\n", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += firstName + "! "
	}
}

func (s *server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("GreetEveryone function was invoked!")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		errors.HandleError("error while receiving req from the client", err)
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{Result: result})
		errors.HandleError("error while sending GreetEveryone response", sendErr)
	}
}

func (s *server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline function was invoked with %v\n", req)

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			fmt.Println("The client canceled the request.")
			return nil, status.Errorf(codes.DeadlineExceeded, "The client canceled the request.")
		}
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetWithDeadlineResponse{
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

	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	errors.HandleError("Failed loading certificates", sslErr)

	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)

	// adding handlers
	greetpb.RegisterGreetServiceServer(s, &server{})

	// start server
	fmt.Println("greet server started, listening on 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
