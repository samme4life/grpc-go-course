package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)

	i := in.Int_1 + in.Int_2

	return &pb.SumResponse{
		Result: fmt.Sprintf("Addition of the input integers %d and %d is: %d", in.Int_1, in.Int_2, i),
	}, nil
}
