package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	loginpb "github.com/sabaruto/streaming-sevice-merger/backend/genproto/authorisation/login/v1"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// TODO: Handle adding services from another function
	authEndpoint := fmt.Sprintf("%s:%s", os.Getenv("AUTHORISATION_SERVICE_HOST"), os.Getenv("AUTHORISATION_SERVICE_PORT"))
	if authEndpoint == "" {
		log.Printf("Cannot find endpoint var, running locally")
		authEndpoint = "localhost:9090"
	}

	err := loginpb.RegisterAuthoriseServiceHandlerFromEndpoint(ctx, mux, authEndpoint, opts)
	if err != nil {
		grpclog.Fatal("failed to register service: %v", err)
	}
	
	log.Printf("server listening at %v", "8081")
	if err = http.ListenAndServe(":8081", mux); err != nil {
		grpclog.Fatal(err)
	}
}