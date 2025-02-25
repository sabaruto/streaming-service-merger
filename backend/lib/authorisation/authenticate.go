package authorisation

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sabaruto/streaming-service-merger/backend/lib/authorisation/postgres/models"
	"github.com/sabaruto/streaming-service-merger/backend/lib/genproto/authenticate/v1"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) GenerateToken(ctx context.Context, request *authenticate.GenerateTokenRequest) (*authenticate.Token, error) {
	log.Printf("Got login request")

	ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Unknown, "error hashing password")
	}

	// Check if the username exists in our database
	customer, err := models.CustomerByNamePassword(ctx, s.db, request.Name, string(passwordHash))
	switch err {
	case sql.ErrNoRows:
		log.Printf("given request %-v", request)
		time.Sleep(1 * time.Second)
		return nil, status.Error(codes.Unauthenticated, "username not found")
	case nil:
		break
	default:
		log.Printf("request data: Name - %s Password - %s", request.Name, request.Password)
		log.Printf("error given: %v", err)
		return nil, status.Errorf(codes.Unknown, "unknown error occurred")
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password))
	if err != nil {
		log.Printf("error given: %v", err)
		return nil, status.Error(codes.Unauthenticated, "username or password does not match")
	}

	// Check if there's already a newStore created
	newStore, err := s.getLatestToken(ctx, customer.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "unknown error occurred")
	}

	return &authenticate.Token{CustomerId: customer.ID.String(), Code: newStore.Code}, nil
}

func (s *server) Authenticate(ctx context.Context, request *authenticate.AuthenticateRequest) (*emptypb.Empty, error) {
	log.Printf("Got authenticate request")

	ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()

	customerID, err := uuid.Parse(request.Token.CustomerId)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	_, err = models.CustomerByID(ctx, s.db, customerID)
	switch err {
	case sql.ErrNoRows:
		return nil, status.Error(codes.Unauthenticated, "username doesn't exist")
	case nil:
		break
	default:
		return nil, status.Error(codes.Unknown, "unknown error occurred")
	}

	store, err := models.TokenStoreByCode(ctx, s.db, request.Token.Code)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "token doesn't exist")
	}

	if customerID != store.CustomerID {
		return nil, status.Error(codes.Unauthenticated, "token id mismatch")
	}

	// Check the token is usable
	if store.ExpireAfter.Compare(time.Now()) < 0 {
		return nil, status.Error(codes.Unauthenticated, "token expired")
	}

	return &emptypb.Empty{}, nil
}

func (s *server) DeleteToken(context.Context, *authenticate.DeleteTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
