package authorisation

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/sabaruto/streaming-service-merger/backend/lib/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	AUTHORISATION_SERVICE_HOST = "localhost:9090"
	GRPC_HOST = "http://localhost:8081"
	DATABASE_URL = "postgres://postgres:postgres@localhost:5432/ssm_authorisation_test"
)

func setup(t *testing.T) (context.Context, *grpc.Server, context.CancelFunc){
	ctx, cancel := context.WithCancel(context.Background())

	lis, err := net.Listen("tcp", AUTHORISATION_SERVICE_HOST)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s, err := NewServer(DATABASE_URL)
	if err != nil {
		t.Errorf("error starting authorisation service: %v", err)
	}

	go func() {
		if err = s.Serve(lis); err != nil {
			t.Errorf("error serving proxy: %v", err)
		}
	}()

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

	return ctx, s, cancel
}

func loginFuncs(input *strings.Reader, expectedCode int) (func(t *testing.T)) {
	return func (t *testing.T) {
		resp, err := http.Post(
			fmt.Sprintf("%s/v1/authorisation/login", GRPC_HOST),
			"application/json",
			input,
		)
	
		if err != nil {
			t.Error("Error connecting to reverse proxy")
		}
	
		t.Logf("Response: %v", resp.StatusCode)

		if expectedCode != resp.StatusCode {
			t.Errorf("Unexpected status code %v", resp.StatusCode)
		}
	}
}

func TestLogin(t *testing.T) {
	_, s, cancel := setup(t)
	defer cancel()
	defer s.Stop()

	time.Sleep(1 * time.Second)

	t.Run("No Input", loginFuncs(&strings.Reader{}, 401))

	r := strings.NewReader(`{"username":"theodosia"}`)
	t.Run("Missing info", loginFuncs(r, 401))

	r = strings.NewReader(`{"username": "theodosia", "password": "theodosia"}`)
	t.Run("No customer", loginFuncs(r, 401))

}