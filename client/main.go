package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	pb "examples/products/products"
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
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	productId := defaultId
	if len(os.Args) > 1 {
		productId, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("No ID passed in: %v", err)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.GetProducts(ctx, &pb.HelloRequest{ProductId: int32(productId)})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product ID: %d, Message: %s", response.GetId(), response.GetMessage())
}
