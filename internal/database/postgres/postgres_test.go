// +build integration

package postgres

import (
	"log"
	"os"
	"testing"
)

const (
	databaseName   = "test_docker"
	migrationsPath = "file://../../../migrations/postgres"
	postgresURL    = "mongodb://docker:example@localhost:8090/test_docker?authSource=admin&ssl=false"
)

func setup() (*Database, error) {
	db, err := New(postgresURL)
	if err != nil {
		return &Database{}, err
	}

	databaseTest := db

	err = databaseTest.Migrate(databaseName, migrationsPath)
	if err != nil {
		return &Database{}, err
	}

	return db, err
}

func TestMain(m *testing.M) {
	db, err := setup()
	defer db.sql.Close()

	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	os.Exit(code)
}
