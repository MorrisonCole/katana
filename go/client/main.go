package main

import (
	"context"
	"log"
	"time"

	pb "com.morrisoncole/katana/go/schema"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDictionaryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	requestDefinition(c, ctx, "かたな")
}

func requestDefinition(client pb.DictionaryClient, context context.Context, word string) {
	r, err := client.RequestDefinition(context, &pb.DefinitionRequest{Word: word})
	if err != nil {
		log.Fatalf("could not get definition: %v", err)
	}
	log.Printf("Got definition: %s", r.Definition)
}
