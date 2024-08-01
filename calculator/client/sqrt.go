package main

import (
	"context"
	"log"

	pb "github.com/ddcad2030/grpc-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt client")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %v\n", e.Message())
			log.Printf("Error message from server: %v\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("We probaly send a negative number: ")
				return
			}

		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}
	log.Printf("Sqrt: %f\n", res.Result)
}
