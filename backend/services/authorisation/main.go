package main

import (
	"flag"
	"log"
	"net"

	loginpb "github.com/sabaruto/streaming-sevice-merger/backend/genproto/authorisation/login/v1"
	"google.golang.org/grpc"
	// Update
)

var authServerEndpoint = flag.String("authorisation-server-endpoint", "localhost:9090", "Auth server endpoint")


func start_server(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	loginpb.RegisterAuthoriseServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	return s.Serve(lis)
}


func main() {
	flag.Parse()

	if err := start_server(*authServerEndpoint); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
