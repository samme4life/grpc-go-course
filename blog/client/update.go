package main

import (
	"context"
	pb "grpc-go-course/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Updated AuthorId",
		Title:    "Updated Title",
		Content:  "Updated Content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error updating blog %v\n", err)
	}
	log.Println("Blog has been updated")
}
