package main

import (
	"context"
	pb "grpc-go-course/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error reading blog %v\n", err)
	}
	log.Printf("Blog read: %s\n", res.Id)
	return res
}
