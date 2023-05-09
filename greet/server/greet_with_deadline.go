package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-go-course/greet/proto"
	"log"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client has cancelled the request!")
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}

		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
