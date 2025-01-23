package main

import (
	"context"
	"fmt"
	"log"
	"net"

	loginpb "github.com/sabaruto/streaming-sevice-merger/backend/genproto/authorisation/login/v1"
	"google.golang.org/grpc"
)

type server struct {
	loginpb.UnimplementedAuthoriseServiceServer
}

func (s *server) Login(ctx context.Context, request *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	log.Printf("Got login request")
	return &loginpb.LoginResponse{Result: &loginpb.LoginResponse_Success{}}, nil
}

func start_server(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
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
