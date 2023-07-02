package main

import (
	"context"
	pb "grpc-go-course/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "samme4life",
		Title:    "How to Secure your gRPC Server?",
		Content:  "Content of the blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Error creating blog %v\n", err)
	}
	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
