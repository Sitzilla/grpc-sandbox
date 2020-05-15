
//go:generate protoc -I games --go_out=plugins=grpc:games games/games.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
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

type game struct {
	Id int32
	Name string
	Company string
	Type string
	ReleaseYear int32
}

var games []game

func (s *server) GetGames(ctx context.Context, in *protobuff.GameRequest) (*protobuff.GameReply, error) {
	log.Printf("Request received for id: %d" , in.GetId())
	mygame, err := getGameById(in.GetId())

	if err != nil {
		log.Printf("Could not find game: %v", err)
	}

	//log.Print(mygame)
	return &protobuff.GameReply{Id: mygame.Id, Name: mygame.Name, Company: mygame.Company, Type: mygame.Type, ReleaseYear: mygame.ReleaseYear}, nil
}

func getGameById(id int32) (game, error) {
	for _, game := range games {
		if game.Id == id {
			return game, nil
		}
	}

	return game{}, fmt.Errorf("Cannot find game with ID: %d", id)
}

func main() {
	game1 := game{1, "Zelda, Breath of the Wild", "Nintendo", "RPG", 2017}
	game2 := game{2, "Starcraft", "Blizzard", "RTS", 1998}
	game3 := game{3, "Stardew Valley", "ConcernedApe", "Farming?", 2016}

	games = append(games, game1)
	games = append(games, game2)
	games = append(games, game3)
	log.Printf("Games initialized")

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
