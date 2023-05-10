package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpc-go-course/blog/proto"
	"log"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Printf("ListBlogs was invoked %v\n", in)

	cur, err := collection.Find(context.Background(), primitive.D{})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("An error occurred while getting the blog list %v\n", err),
		)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			fmt.Println("An error occurred while closing the cursor ", err)
		}
	}(cur, context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("An error occurred while decoding data from MongoDB: %v\n", err),
			)
		}
		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("An error occurred while iterating the blog list: %v\n", err),
		)
	}

	return nil

}
