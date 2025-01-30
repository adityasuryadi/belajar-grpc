package main

import (
	"context"
	"log"
	"time"

	pb "github.com/adityasuryadi/belajar-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreaterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Printf("an error occured: %v", err)
		return
	}
	log.Printf("Message from server: %s", r.Message)
}
