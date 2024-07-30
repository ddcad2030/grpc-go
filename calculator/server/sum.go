package main

import (
	"log"

	pb "github.com/ddcad2030/grpc-go/calculator/proto"

	"context"
)

func (s *Server) Sum(c context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function invoked %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
