package main

import (
	"context"
	"log"

	pb "github.com/ddcad2030/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("create blog was invoked")
	blog := &pb.Blog{
		AuthorId: "David",
		Title:    "First Blog",
		Content:  "Content of first blog",
	}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res.Id)
	return res.Id
}
