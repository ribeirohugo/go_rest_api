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
	testId    = "1"
	testEmail = "email@domain"
	testName  = "name"

	testEmailUpdate = "mail@domain"
	testNameUpdate  = "name surname"
)

var testUser = model.User{
	Id:    testId,
	Email: testEmail,
	Name:  testName,
}

func TestDatabase_CRUD(t *testing.T) {
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

	user.Name = testNameUpdate
	user.Email = testEmailUpdate

	err = databaseForTest.UpdateUser(context.Background(), user)
	require.NoError(t, err)

	updatedUser, err := databaseForTest.FindUser(context.Background(), userId)
	require.NoError(t, err)

	assert.Equal(t, testNameUpdate, updatedUser.Name)
	assert.Equal(t, testEmailUpdate, updatedUser.Email)

	err = databaseForTest.DeleteUser(context.Background(), userId)
	require.NoError(t, err)

	user, err = databaseForTest.FindUser(context.Background(), userId)

	require.Error(t, err)
}
