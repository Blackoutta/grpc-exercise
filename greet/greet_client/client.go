package main

import (
	"context"
	"fmt"
	"learnings/grpc/greet/greetpb"
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
	c := greetpb.NewGreetServiceClient(cc)

	// invoking grpc calls
	doUnary(c)
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
	log.Printf("Response from Greet: %v", res.Result)
}
