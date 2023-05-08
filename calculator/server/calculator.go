package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-go-course/calculator/proto"
)

func (s Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)

	i := in.Int_1 + in.Int_2

	return &pb.CalculatorResponse{
		Result: fmt.Sprintf("Addition of the input integers %d and %d is: %d", in.Int_1, in.Int_2, i),
	}, nil
}
