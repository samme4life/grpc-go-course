package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"log"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Println("doCalculate was invoked...")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Int_1: 45,
		Int_2: -87,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("Result: %s\n", res.Result)

}
