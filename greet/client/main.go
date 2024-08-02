package main

import (
	"log"
	"time"

	pb "github.com/ddcad2030/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	// doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// DoGreetEveryone(c)
	doGreetWithDeadline(c, 5*time.Second)
	// doGreetWithDeadline(c, 1 *time.Second) // server waiting 3second, exceecded before 3s error
}
