package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	gw "github.com/sabaruto/streaming-sevice-merger/backend/genproto/authorisation/login/v1" // Update
)

var grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")

func start_proxy() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterAuthoriseServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		grpclog.Info("failed to register service: %v", err)
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("server listening at %v", "8081")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	log.Printf("Starting GRPC Proxy server")
	if err := start_proxy(); err != nil {
		grpclog.Fatal(err)
	}
}