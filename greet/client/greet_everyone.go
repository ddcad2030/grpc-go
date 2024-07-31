package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ddcad2030/grpc-go/greet/proto"
)

func DoGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "David"},
		{FirstName: "Tuan"},
		{FirstName: "Loc"},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Receiving: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
