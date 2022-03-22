package postgres

import (
	"context"
	"github.com/ribeirohugo/golang_startup/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var databaseTest Database

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
	userId, err := databaseTest.CreateUser(context.Background(), testUser)
	assert.NoError(t, err)

	user, err := databaseTest.FindUser(context.Background(), userId)
	assert.NoError(t, err)

	assert.Equal(t, testEmail, user.Email)
	assert.Equal(t, testName, user.Name)
}
