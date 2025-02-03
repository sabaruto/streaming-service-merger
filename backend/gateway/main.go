package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sabaruto/streaming-service-merger/backend/lib/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	authEndpoint := fmt.Sprintf("%s:%s", os.Getenv("AUTHORISATION_SERVICE_HOST"), os.Getenv("AUTHORISATION_SERVICE_PORT"))
	mux, err := gateway.StartReverseProxy(ctx, opts, authEndpoint)
	if err == nil {
		grpclog.Fatal(err)
	}
	
	log.Printf("server listening at %v", "8081")
	if err = http.ListenAndServe(":8081", mux); err != nil {
		grpclog.Fatal(err)
	}
}
