package authorisation

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/sabaruto/streaming-service-merger/backend/lib/authorisation/postgres/models"
	"github.com/sabaruto/streaming-service-merger/backend/lib/genproto/customer/v1"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCustomer(name string, password string) (customer *models.Customer, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		err = status.Error(codes.Unknown, "error hashing password")
		return
	}

	customer = &models.Customer{
		ID:       uuid.New(),
		Name:     name,
		Password: string(passwordHash),
	}
	return
}

func (s *server) CreateCustomer(ctx context.Context, request *customer.CreateCustomerRequest) (empty *emptypb.Empty, err error) {
	log.Printf("Got sign-up request")

	ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()

	newUser, err := NewCustomer(request.Name, request.Password)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "username already exists")
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
		return nil, status.Error(codes.Unknown, "unknown error occurred")
	}

	return &emptypb.Empty{}, nil
}
func (s *server) GetCustomer(ctx context.Context, request *customer.GetCustomerRequest) (*customer.Customer, error) {
	id, err := uuid.Parse(request.CustomerId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "cannot parse uuid")
	}

	modelCustomer, err := models.CustomerByID(ctx, s.db, id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "customer not found")
	}

	responseCustomer := &customer.Customer{
		Name:     modelCustomer.Name,
		Password: modelCustomer.Password,
	}
	return responseCustomer, nil
}
func (s *server) UpdateCustomer(ctx context.Context, request *customer.UpdateCustomerRequest) (*customer.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (s *server) DeleteCustomer(context.Context, *customer.DeleteCustomerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
