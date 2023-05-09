package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-go-course/greet/proto"
	"log"
)

var addr string = "localhost:50051"

func main() {
	opts := []grpc.DialOption{}
	tls := true //change this to false if needed
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error loading CA trust certificate %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr, opts...)

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

	doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	//doGreetEveryone(c)
	//doGreetWithDeadline(c, 1*time.Second)
}
