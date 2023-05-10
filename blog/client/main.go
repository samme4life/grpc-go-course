package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-course/blog/proto"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error at closing the connection: %v\n", err)
		}
	}(conn)

	c := pb.NewBlogServiceClient(conn)

	//doCalculate(c)
	id := createBlog(c)
	readBlog(c, id) //valid id
	//readBlog(c, "aNonExistingId") // invalid id
	updateBlog(c, id)
	listBlogs(c)
}
