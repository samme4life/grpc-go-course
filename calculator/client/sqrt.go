package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-go-course/calculator/proto"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if codes.InvalidArgument == e.Code() {
				log.Println("We probably sent a negative number")
			}
		} else {
			log.Fatalf("A non gRPC error occurred: %v\n")
		}
		return
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
