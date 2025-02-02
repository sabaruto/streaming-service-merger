package main

import (
	"fmt"
	"log"
	"net"
	"os"

	loginpb "github.com/sabaruto/streaming-sevice-merger/backend/genproto/authorisation/login/v1"
	"google.golang.org/grpc"
)

func main() {
	endpoint := fmt.Sprintf(":%s",  os.Getenv("AUTHORISATION_SERVICE_PORT"))

	if endpoint == "" {
		log.Printf("Cannot find endpoint vars, serving localhost")
		endpoint = "localhost:9090"
	}

	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	loginpb.RegisterAuthoriseServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
