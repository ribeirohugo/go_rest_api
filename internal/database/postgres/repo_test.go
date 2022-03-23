package postgres

import (
	"context"
	"testing"

	"github.com/ribeirohugo/golang_startup/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
)

const (
	testId    = "00000000-0000-0000-0000-000000000000"
	testEmail = "email@domain"
	testName  = "name"
)

var testUser = model.User{
	Id:    testId,
	Email: testEmail,
	Name:  testName,
}

func TestDatabase_CreateUser(t *testing.T) {
	container, err := setup(t)
	defer shutdown(t, container)
	require.NoError(t, err)

	databaseForTest := buildClient(t, container)

	userId, err := databaseForTest.CreateUser(context.Background(), testUser)
	require.NoError(t, err)

	user, err := databaseForTest.FindUser(context.Background(), userId)
	require.NoError(t, err)

	assert.Equal(t, testEmail, user.Email)
	assert.Equal(t, testName, user.Name)
}
