package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/ddcad2030/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"
var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blog_db").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
