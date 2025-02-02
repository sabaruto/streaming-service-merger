package main

import (
	"context"
	"log"

	authpb "github.com/sabaruto/streaming-service-merger/backend/genproto/v1/authorisation"
	commonpb "github.com/sabaruto/streaming-service-merger/backend/genproto/v1/common"
	"github.com/xo/dburl"
)

type server struct {
	authpb.UnimplementedAuthoriseServiceServer
	dbURL *dburl.URL
}

func (s *server) SignUp(ctx context.Context, request *authpb.SignUpRequest) (*commonpb.AuthenticateResponse, error) {
	log.Printf("Got sign-up request")

	// Add new entry to database

	// Return appropriately

	return &commonpb.AuthenticateResponse{Result: &commonpb.AuthenticateResponse_Success{}}, nil
}

func (s *server) Login(ctx context.Context, request *authpb.LoginRequest) (*commonpb.AuthenticateResponse, error) {
	log.Printf("Got login request")

	// Check if the username exists in our database
	
	// Check if hashed password matches

	// If above are true, return true, else raise error as expected
	return &commonpb.AuthenticateResponse{Result: &commonpb.AuthenticateResponse_Success{}}, nil
}
