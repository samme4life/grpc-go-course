package main

import (
	"context"
	"fmt"
	"github.com/samme4life/grpc-go-course/calculate/calculatepb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	calculatepb.UnimplementedCalculateServiceServer
}

func (*server) Calculate(ctx context.Context, req *calculatepb.CalculateRequest) (*calculatepb.CalculateResponse, error) {
	fmt.Printf("Calculate Function was invoked with %v\n", req)

	op1 := req.GetOperands().GetO1()
	op2 := req.GetOperands().GetO2()

	result := op1 + op2

	res := &calculatepb.CalculateResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Server Started!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatepb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
