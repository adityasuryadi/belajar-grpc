package main

import (
	"context"
	"log"
	"net"

	pb "github.com/adityasuryadi/belajar-grpc/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreaterServer
}

func (s *server) SayHelo(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("Received: ", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreaterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
