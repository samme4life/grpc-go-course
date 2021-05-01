package main

import (
	"context"
	"fmt"
	"github.com/samme4life/grpc-go-course/calculate/calculatepb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello, I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	//c := greetpb.NewGreetServiceClient(cc)
	c := calculatepb.NewCalculateServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatepb.CalculateServiceClient) {
	fmt.Println("Starting to do Unary RPC...")

	req := &calculatepb.CalculateRequest{
		Operands: &calculatepb.Operands{
			O1: 42,
			O2: 65,
		},
	}
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from the Calculate RPC: %v", res.Result)
}
