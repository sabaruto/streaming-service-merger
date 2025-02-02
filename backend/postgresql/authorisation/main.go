package main

import (
	"context"
	"log"
	"os/user"

	"github.com/sabaruto/streaming-sevice-merger/backend/postgresql/authorisation/models"
	"github.com/xo/dburl"
	"github.com/xo/dburl/passfile"
)

// TODO:
// - Move to services file
// - Allow connection from different placies .i.e. kubectl + localhost
// - Add functions to Add, remove and query different customers
func main()  {
	ctx := context.Background()
	pwd, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	url, err := dburl.Parse("postgres://admin:password@localhost")
	if err != nil {
		log.Fatal(err)
	}

	db, err := passfile.OpenURL(url, pwd.HomeDir, "xopass")
	if err != nil {
		log.Fatal(err)
	}

	theodosia := models.Customer{
		Name: "Theodosia",
		Password: "password",
	}

	if err := theodosia.Save(ctx, db); err != nil {
		log.Fatal(err)
	}
}