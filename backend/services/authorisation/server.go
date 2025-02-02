package main

import (
	"context"
	"log"

	loginpb "github.com/sabaruto/streaming-sevice-merger/backend/genproto/authorisation/login/v1"
)

type server struct {
	loginpb.UnimplementedAuthoriseServiceServer
}

func (s *server) Login(ctx context.Context, request *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	log.Printf("Got login request")

	// Check if the username exists in our database
	
	// Check if hashed password matches

	// If above are true, return true, else raise error as expected
	return &loginpb.LoginResponse{Result: &loginpb.LoginResponse_Success{}}, nil
}
