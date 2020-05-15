package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	protobuff "examples/games/games"
	"google.golang.org/grpc"
)

const (
	address   = "localhost:50051"
	defaultId = 1
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protobuff.NewGameServiceClient(conn)

	// Contact the server and print out its response.
	gameId := defaultId
	if len(os.Args) > 1 {
		gameId, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("No ID passed in: %v", err)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.GetGames(ctx, &protobuff.GameRequest{Id: int32(gameId)})
	if err != nil {
		log.Fatalf("Could not get game: %v", err)
	}
	log.Printf("Game ID: %d, Name: %s, Company: %s, Type: %s, Release Year: %d", response.GetId(), response.GetName(), response.GetCompany(), response.GetType(), response.GetReleaseYear())
}
