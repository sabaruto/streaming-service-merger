package authorisation

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sabaruto/streaming-service-merger/backend/lib/genproto/authenticate/v1"
	"github.com/sabaruto/streaming-service-merger/backend/lib/genproto/customer/v1"
	"github.com/xo/dburl"
	"google.golang.org/grpc"
)

const TIMEOUT = 5 * time.Minute

type server struct {
	authenticate.UnimplementedAuthenticateServiceServer
	customer.UnimplementedCustomerServiceServer
	db *sql.DB
}

func StartService(url string) (s *grpc.Server, err error) {
	db, err := dburl.Open(url)
	if err != nil {
		err = fmt.Errorf("error parsing url: %v", err)
		return
	}

	s = grpc.NewServer()
	authenticate.RegisterAuthenticateServiceServer(s, &server{db: db})
	customer.RegisterCustomerServiceServer(s, &server{db: db})
	return
}
