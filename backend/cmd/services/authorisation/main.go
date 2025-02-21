package main

import (
	"fmt"
	"net"
	"os"

	"github.com/sabaruto/streaming-service-merger/backend/lib/authorisation"
	log "github.com/sirupsen/logrus"
)

func main() {
	endpoint := fmt.Sprintf(":%s", os.Getenv("AUTHORISATION_SERVICE_PORT"))
	url := os.Getenv("AUTHORISATION_DATABASE_URL")
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s, err := authorisation.StartService(url)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
