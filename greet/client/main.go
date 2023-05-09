package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-course/greet/proto"
	"log"
	"time"
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

	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	//doGreetEveryone(c)
	doGreetWithDeadline(c, 1*time.Second)
}
