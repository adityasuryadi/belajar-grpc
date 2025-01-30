package main

import (
	"context"
	"log"
	"net"

	pb "github.com/adityasuryadi/belajar-grpc/proto"
	"google.golang.org/grpc"
)

type server struct {
	// mengimplementasikan server
	pb.UnimplementedGreaterServer
}

// method say hello dari protobuff
// implemetasi dari server
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// req.Getname() hasil dari generate protobuff
	log.Println("Received: ", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// buat server grpc
	s := grpc.NewServer()

	// menjalankan server grpc menggunakan struct server di atas
	pb.RegisterGreaterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
