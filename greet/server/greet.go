package main

import (
	"context"
	"log"

	pb "github.com/ddcad2030/grpc-go/greet/proto"
)

func (s *Server) Greet(c context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
