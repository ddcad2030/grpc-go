package main

import (
	"context"
	"log"

	pb "github.com/ddcad2030/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("update blog was invoked")

	data := &pb.Blog{
		Id:       id,
		AuthorId: "Not david",
		Title:    "update blog",
		Content:  "update content",
	}

	_, err := c.UpdateBlog(context.Background(), data)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Printf("Blog was updated")
}
