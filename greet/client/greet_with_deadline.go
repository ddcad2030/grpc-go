package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ddcad2030/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, t time.Duration) {
	log.Println("doGreetWithDeadline was invoked")
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	req := &pb.GreetRequest{
		FirstName: "David",
	}
	res, err := c.GreetDeadLine(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded !")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}

		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}

	}
	log.Printf("GreetWithDeadline: %v\n", res.Result)
}
