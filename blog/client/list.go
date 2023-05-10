package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpc-go-course/blog/proto"
	"io"
	"log"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("listBlogs was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error occurred while calling the ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error receiving Blog List: %v\n", err)
		}

		fmt.Println(res)
	}
}
