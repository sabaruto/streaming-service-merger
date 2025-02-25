package authorisation

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/sabaruto/streaming-service-merger/backend/lib/gateway"
	"github.com/sabaruto/streaming-service-merger/backend/lib/genproto/customer/v1"
	"github.com/stretchr/testify/assert"
	"github.com/xo/dburl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	AUTHORISATION_SERVICE_HOST = "localhost:9090"
	GRPC_HOST                  = "http://localhost:8081"
	DATABASE_URL               = "postgres://postgres:postgres@localhost:5432/ssm_authorisation_test"
)

func setupDB(t *testing.T) (*sql.Tx, *grpc.Server, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	// Setup authorisation service
	lis, err := net.Listen("tcp", AUTHORISATION_SERVICE_HOST)
	if err != nil {
		t.Errorf("failed to listen: %v", err)
	}

	s, err := StartService(DATABASE_URL)
	if err != nil {
		t.Errorf("error starting authorisation service: %v", err)
	}

	go func() {
		if err = s.Serve(lis); err != nil {
			t.Errorf("error serving proxy: %v", err)
		}
	}()

	// Setup grpc gateway
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	mux, err := gateway.StartReverseProxy(ctx, opts, AUTHORISATION_SERVICE_HOST)
	if err != nil {
		t.Errorf("error starting reverse proxy: %v", err)
	}

	go func() {
		if err = http.ListenAndServe(":8081", mux); err != nil {
			t.Errorf("error serving proxy: %v", err)
		}
	}()

	db, err := dburl.Open(DATABASE_URL)
	if err != nil {
		t.Errorf("error connecting to db: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Errorf("error starting transaction %v", err)
	}

	return tx, s, cancel
}

func teardown(t *testing.T) {

}

func loginFuncs(body string, expectedCode int) func(t *testing.T) {
	input := strings.NewReader(body)
	return func(t *testing.T) {
		resp, err := http.Post(
			fmt.Sprintf("%s/v1/authorisation/login", GRPC_HOST),
			"application/json",
			input,
		)

		if err != nil {
			t.Error("Error connecting to reverse proxy")
		}

		if expectedCode != resp.StatusCode {
			t.Log(resp.Header["Www-Authenticate"])
			t.Errorf("Unexpected status code %v", resp.StatusCode)
		}
	}
}

// func TestLogin(t *testing.T) {
// 	ctx := context.Background()

// 	tx, s, cancel := setupDB(t)
// 	defer cancel()
// 	defer s.Stop()

// 	testUser, err := NewCustomer("Test", "Test")
// 	if err != nil {
// 		t.Errorf("error creating new user %v", err)
// 	}

// 	err = testUser.Upsert(ctx, tx)
// 	if err != nil {
// 		t.Errorf("error saving new user %v", err)
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		t.Errorf("error commiting new user %v", err)
// 	}

// 	t.Run("No Input", loginFuncs("", 401))
// 	t.Run("Missing info", loginFuncs(`{"username":"theodosia"}`, 401))
// 	t.Run("Customer does not exist", loginFuncs(`{"username": "theodosia", "password": "theodosia"}`, 401))
// 	t.Run("Correct Info", loginFuncs(`{"username": "Test", "password": "Test"}`, 200))

// 	tx.Rollback()
// }

func TestCustomer(t *testing.T) {
	var (
		db  *sql.DB
		err error
	)

	ctx := context.Background()

	if db, err = dburl.Open(DATABASE_URL); err != nil {
		t.Errorf("Error opening database: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Errorf("error starting transaction %v", err)
	}

	testUser, err := NewCustomer("Test", "Test")
	if err != nil {
		t.Errorf("error creating new user %v", err)
	}

	err = testUser.Upsert(ctx, tx)
	if err != nil {
		t.Errorf("error saving new user %v", err)
	}

	err = tx.Commit()
	if err != nil {
		t.Errorf("error committing new user %v", err)
	}

	s := &server{db: db}

	t.Run("GetCustomerEmptyResponse", func(t *testing.T) {
		request := &customer.GetCustomerRequest{}

		_, err := s.GetCustomer(ctx, request)

		assert.NotNil(t, err)
		assert.Equal(t, err, status.Error(codes.InvalidArgument, "cannot parse uuid"))
	})

	t.Run("GetCustomerWrongFormat", func(t *testing.T) {
		request := &customer.GetCustomerRequest{
			CustomerId: "wrong",
		}

		_, err := s.GetCustomer(ctx, request)

		assert.NotNil(t, err)
		assert.Equal(t, err, status.Error(codes.InvalidArgument, "cannot parse uuid"))
	})

	t.Run("GetCustomerNoUser", func(t *testing.T) {
		request := &customer.GetCustomerRequest{
			CustomerId: uuid.New().String(),
		}

		_, err := s.GetCustomer(ctx, request)

		assert.NotNil(t, err)
		assert.Equal(t, err, status.Error(codes.InvalidArgument, "customer not found"))
	})

	t.Run("GetCustomerNoUser", func(t *testing.T) {
		request := &customer.GetCustomerRequest{
			CustomerId: uuid.New().String(),
		}

		_, err := s.GetCustomer(ctx, request)

		assert.NotNilf(t, err, "error getting customer:")
		assert.Equal(t, err, status.Error(codes.InvalidArgument, "customer not found"))
	})

	t.Run("GetCustomerPass", func(t *testing.T) {
		request := &customer.GetCustomerRequest{
			CustomerId: uuid.New().String(),
		}

		_, err := s.GetCustomer(ctx, request)

		assert.NotNilf(t, err, "error getting customer:")
		assert.Equal(t, err, status.Error(codes.InvalidArgument, "customer not found"))
	})

	t.Run("UpdateCustomer", func(t *testing.T) {
		request := &customer.UpdateCustomerRequest{
			Customer: &customer.Customer{
				Name:     "test",
				Password: "test",
			},
		}
		_, err := s.UpdateCustomer(ctx, request)
		if err != nil {
			t.Errorf("error getting customer: %v", err)
		}
	})

	t.Run("DeleteCustomer", func(t *testing.T) {
		request := &customer.DeleteCustomerRequest{
			CustomerId: "test",
		}
		_, err := s.DeleteCustomer(ctx, request)
		if err != nil {
			t.Errorf("error getting customer: %v", err)
		}
	})
}
