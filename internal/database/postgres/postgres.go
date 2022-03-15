package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const postgresDriveName = "postgres"

type SQL interface {
	Ping() error
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// Database represents an initialised client to the database.
type Database struct {
	sql SQL
}

func New(address string) (*Database, error) {
	db, err := sql.Open(postgresDriveName, address)
	if err != nil {
		return &Database{}, err
	}

	err = db.Ping()
	if err != nil {
		return &Database{}, fmt.Errorf("error pinging database connection: %s", err.Error())
	}

	return &Database{sql: db}, nil
}
