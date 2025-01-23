package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	gw "github.com/sabaruto/streaming-sevice-merger/backend/genproto/authorisation/login/v1" // Update
)

func start_proxy() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterAuthoriseServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		grpclog.Info("fiailed to register service: %v", err)
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	grpclog.Infof("server listening at %v", "8081")
	return http.ListenAndServe(":8081", mux)
}
