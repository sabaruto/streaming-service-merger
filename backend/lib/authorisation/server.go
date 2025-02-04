package authorisation

import (
	"context"

	"database/sql"

	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sabaruto/streaming-service-merger/backend/lib/authorisation/postgres/models"
	authpb "github.com/sabaruto/streaming-service-merger/backend/lib/genproto/v1/authorisation"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	authpb.UnimplementedAuthoriseServiceServer
	db *sql.DB
}

// TODO: Fix invalid UUID length
const TIMEOUT = 5 * time.Minute

func (s *server) SignUp(ctx context.Context, request *authpb.CredsRequest) (*emptypb.Empty, error) {
	log.Printf("Got sign-up request")

	ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()

	_, err := models.CustomerByName(ctx, s.db, request.Username)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "username already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Unknown, "error hashing password")
	}
	
	newUser := &models.Customer {
		ID: uuid.New(),
		Name: request.Username,
		Password: string(passwordHash),
	}

	err = newUser.Save(ctx, s.db)

	switch err {
	case models.ErrAlreadyExists:
		return nil, status.Error(codes.AlreadyExists, "username already exists")
	case context.DeadlineExceeded:
		return nil, status.Error(codes.DeadlineExceeded, "timeout error")
	case nil:
		break
	default:
		return nil, status.Error(codes.Unknown, "unkown error occured")
	}

	return &emptypb.Empty{}, nil
}

func (s *server) Login(ctx context.Context, request *authpb.CredsRequest) (*authpb.Auth, error) {
	log.Printf("Got login request")

	ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()

	// Check if the username exists in our database
	customer, err := models.CustomerByName(ctx, s.db, request.Username)
	switch err {
	case sql.ErrNoRows:
		return nil, status.Error(codes.Unauthenticated, "username or password does not match")
	case nil:
		break
	default:
		log.Print("request data", request.Username, request.Password)
		return nil, status.Errorf(codes.Unknown, "unkown error occured")
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "username or password does not match")
	}

	// Check if there's already a newStore created
	newStore, err := s.GetLatestToken(ctx, customer.ID)
	if err != nil {
		return nil, status.Error(codes.Unknown, "unkown error occured")
	}

	return &authpb.Auth{ClientID: customer.ID.String(), AuthCode: newStore.Token}, nil
}

func (s *server) Authenticate(ctx context.Context, request *authpb.Auth) (*emptypb.Empty, error) {
	log.Printf("Got authenticate request")

	ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()

	customerID, err := uuid.Parse(request.ClientID)
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
		return nil, status.Error(codes.Unknown, "unkown error occured")
	}

	store, err := models.TokenStoreByToken(ctx, s.db, request.AuthCode)
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