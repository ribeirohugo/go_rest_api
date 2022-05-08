// +build integration

package mysql

import (
	"context"
	"testing"

	"github.com/ribeirohugo/golang_startup/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testID    = "1"
	testEmail = "email@domain"
	testName  = "name"

	testEmailUpdate = "mail@domain"
	testNameUpdate  = "name surname"
)

var testUser = model.User{
	ID:    testID,
	Email: testEmail,
	Name:  testName,
}

func TestDatabase_CRUD(t *testing.T) {
	container, err := setup(t)
	defer shutdown(t, container)
	require.NoError(t, err)

	databaseForTest := buildClient(t, container)

	userID, err := databaseForTest.CreateUser(context.Background(), testUser)
	require.NoError(t, err)

	user, err := databaseForTest.FindUser(context.Background(), userID)
	require.NoError(t, err)

	assert.Equal(t, testEmail, user.Email)
	assert.Equal(t, testName, user.Name)

	user.Name = testNameUpdate
	user.Email = testEmailUpdate

	// FindAll
	users2, err := databaseForTest.FindAllUsers(context.Background(), 0, 20)
	require.NoError(t, err)

	assert.Len(t, users2, 1)

	users2, err = databaseForTest.FindAllUsers(context.Background(), 1, 0)
	require.NoError(t, err)

	assert.Empty(t, users2)

	err = databaseForTest.UpdateUser(context.Background(), user)
	require.NoError(t, err)

	updatedUser, err := databaseForTest.FindUser(context.Background(), userID)
	require.NoError(t, err)

	assert.Equal(t, testNameUpdate, updatedUser.Name)
	assert.Equal(t, testEmailUpdate, updatedUser.Email)

	err = databaseForTest.DeleteUser(context.Background(), userID)
	require.NoError(t, err)

	user, err = databaseForTest.FindUser(context.Background(), userID)

	require.Error(t, err)
}
