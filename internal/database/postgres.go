package database

import (
	"database/sql"
)

// Postgres represents an initialised client to the database.
type Postgres struct {
	*sql.DB
}

func NewPostgres(address string) (*Postgres, error) {
	// Connect to database, run migrations, etc.
	return &Postgres{DB: &sql.DB{}}, nil
}
