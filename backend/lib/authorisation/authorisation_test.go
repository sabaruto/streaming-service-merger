package authorisation

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/sabaruto/streaming-service-merger/backend/lib/authorisation/postgres/models"
	"github.com/sabaruto/streaming-service-merger/backend/lib/genproto/customer/v1"
	"github.com/stretchr/testify/assert"
	"github.com/xo/dburl"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	AUTHORISATION_SERVICE_HOST = "localhost:9090"
	GRPC_HOST                  = "http://localhost:8081"
	DATABASE_URL               = "postgres://postgres:postgres@localhost:5432/ssm_authorisation_test"
)

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
// 		t.Errorf("error committing new user %v", err)
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

	staticCustomer, err := NewCustomer("Test", "Test")
	if err != nil {
		t.Errorf("error creating new user %v", err)
	}

	err = staticCustomer.Upsert(ctx, tx)
	if err != nil {
		t.Errorf("error saving new user %v", err)
	}

	updateCustomer, err := NewCustomer("update", "update")
	if err != nil {
		t.Errorf("error creating static user %v", err)
	}

	err = updateCustomer.Upsert(ctx, tx)
	if err != nil {
		t.Errorf("error saving static user %v", err)
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

	t.Run("GetCustomerPass", func(t *testing.T) {
		request := &customer.GetCustomerRequest{
			CustomerId: staticCustomer.ID.String(),
		}

		response, err := s.GetCustomer(ctx, request)

		assert.Nilf(t, err, "error getting customer: %v", err)
		assert.Equal(t, response.Name, "Test")
		assert.NotNil(t, bcrypt.CompareHashAndPassword([]byte(response.Password), []byte("test")))
	})

	t.Run("UpdateCustomer", func(t *testing.T) {
		request := &customer.UpdateCustomerRequest{
			Customer: &customer.Customer{
				Name:     "test",
				Password: "test",
			},
		}
		response, err := s.UpdateCustomer(ctx, request)

		assert.Nilf(t, err, "error getting customer: %v", err)

		if response == nil {
			t.FailNow()
		}

		assert.Equal(t, response.Name, "test")
		assert.NotNilf(t, bcrypt.CompareHashAndPassword([]byte(response.Password), []byte("test")), "password not updated")
	})

	t.Run("DeleteCustomer", func(t *testing.T) {
		id := ""
		request := &customer.DeleteCustomerRequest{
			CustomerId: id,
		}
		response, err := s.DeleteCustomer(ctx, request)
		if err != nil {
			t.Errorf("error getting customer: %v", err)
		}

		assert.Nilf(t, err, "error getting customer: %v", err)

		if response == nil {
			t.FailNow()
		}

		customer, err := models.CustomerByID(ctx, db, uuid.MustParse(id))
		assert.Nil(t, customer)
		assert.NotNil(t, err)
	})
}
