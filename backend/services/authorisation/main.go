package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/sabaruto/streaming-service-merger/backend/genproto/v1/authorisation"
	"github.com/xo/dburl"
	"google.golang.org/grpc"
)

func main() {
	endpoint := fmt.Sprintf(":%s",  os.Getenv("AUTHORISATION_SERVICE_PORT"))
	url := os.Getenv("DATABASE_URL")

	if endpoint == "" {
		log.Printf("Cannot find endpoint env, serving localhost")
		endpoint = "localhost:9090"
	}

	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dbURL, err := dburl.Parse(url)
	if err != nil {
		log.Fatalf("error pasing url: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthoriseServiceServer(s, &server{dbURL: dbURL})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
