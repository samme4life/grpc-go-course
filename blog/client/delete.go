package main

import (
	"context"
	pb "grpc-go-course/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")

	req := &pb.BlogId{Id: id}

	_, err := c.DeleteBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error deleting blog %v\n", err)
	}
	log.Printf("Blog was deleted")
}
