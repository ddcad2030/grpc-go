package main

import (
	"context"
	"log"

	pb "github.com/ddcad2030/grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do Greet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "David",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %v\n", res.Result)
}
