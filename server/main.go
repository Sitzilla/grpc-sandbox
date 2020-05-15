
//go:generate protoc -I ../products --go_out=plugins=grpc:../products ../products/products.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	pb "examples/products/products"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) GetProducts(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Request received for id: %d" , in.GetProductId())
	return &pb.HelloReply{Id: 42, Message: "Zelda: BOTW"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("Server started on port: %s", port)
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
