package postgres

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	postgresMigration "github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/lib/pq"
)

const (
	postgresDriveName = "postgres"
	schemaName        = "schema_migrations"
)

// Database represents an initialised client to the database.
type Database struct {
	sql *sql.DB
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

func (db *Database) Migrate(databaseName string, migrationsPath string) error {

	postgresConfig := postgresMigration.Config{
		SchemaName: schemaName,
	}

	driver, err := postgresMigration.WithInstance(db.sql, &postgresConfig)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, databaseName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		return nil
	}

	return err
}
