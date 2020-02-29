package server

import (
	"context"
	"log"
	"net"

	pb "com.morrisoncole/katana/go/schema"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) RequestDefinition(ctx context.Context, in *pb.DefinitionRequest) (*pb.DefinitionResponse, error) {
	log.Printf("Received definition request: %v", in.Word)

	reading, err := DB.QueryByReading(in.Word)

	if err != nil {
		return &pb.DefinitionResponse{Definition: []string{"No definition found"}}, nil
	}

	return &pb.DefinitionResponse{Definition: reading[0].Definitions}, nil
}

func Serve() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDictionaryServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
