package authorisation

import (
	"fmt"
	"os/user"

	pb "github.com/sabaruto/streaming-service-merger/backend/genproto/v1/authorisation"
	"github.com/xo/dburl/passfile"
	"google.golang.org/grpc"
)

func NewServer(dburl string) (*grpc.Server, error) {
	pwd, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("error getting current user: %v", err)
	}

	db, err := passfile.Open(dburl, pwd.HomeDir, "xopass")
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthoriseServiceServer(s, &server{db: db})

	return s, nil
}