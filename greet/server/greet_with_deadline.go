package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ddcad2030/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetDeadLine(c context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function invoked with %v\n", in)

	for i := 0; i < 3; i++ {
		if c.Err() == context.DeadlineExceeded {
			log.Printf("The client canceled the request !")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
