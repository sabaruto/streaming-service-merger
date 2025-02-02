package database

import (
	"database/sql"
	"fmt"
	"os/user"

	"github.com/xo/dburl"
	"github.com/xo/dburl/passfile"
)

func Connect(url *dburl.URL) (*sql.DB, error) {
	pwd, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("error getting current user: %v", err)
	}

	db, err := passfile.OpenURL(url, pwd.HomeDir, "xopass")
	if err != nil {
		return nil, fmt.Errorf("error opening db: %v", err)
	}

	return db, nil
}