package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating the stream %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Samme"},
		{FirstName: "Ruz"},
		{FirstName: "Luka"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error sending req: %v\n", err)
				return
			}
			time.Sleep(1 * time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Error at CloseSend the stream: %v\n", err)
			return
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while reading the stream: %v\n", err)
				break
			}

			log.Printf("Received: %s\n", res.Result)

		}
		close(waitc)
	}()

	<-waitc

}
