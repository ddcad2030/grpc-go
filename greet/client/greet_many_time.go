package main

import (
	"context"
	"log"

	pb "github.com/ddcad2030/grpc-go/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreet manytime was invoked")
	req := &pb.GreetRequest{
		FirstName: "David",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatal("Error calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}
		if err != nil {
			log.Fatal("Error while stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
