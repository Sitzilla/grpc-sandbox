
//go:generate protoc -I games --go_out=plugins=grpc:games games/games.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	protobuff "examples/games/games"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	protobuff.UnimplementedGameServiceServer
}

func (s *server) GetGames(ctx context.Context, in *protobuff.GameRequest) (*protobuff.GameReply, error) {
	log.Printf("Request received for id: %d" , in.GetId())
	return &protobuff.GameReply{Id: 42, Message: "Zelda: BOTW"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("Server started on port: %s", port)
	protobuff.RegisterGameServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
