package main

import (
	"flag"

	"google.golang.org/grpc/grpclog"
	// Update
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func main() {
	flag.Parse()

	go start_server(9090)

	grpclog.Info("Starting GRPC Proxy server")
	if err := start_proxy(); err != nil {
		grpclog.Fatal(err)
	}
}
