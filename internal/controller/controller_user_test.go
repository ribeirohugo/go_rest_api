package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ribeirohugo/golang_startup/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	idTest    = "00000000-0000-0000-0000-000000000000"
	emailTest = "email@domain"
	nameTest  = "Test"

	jsonOutput = "{\"id\":\"00000000-0000-0000-0000-000000000000\",\"name\":\"Test\",\"email\":\"email@domain\"}\n"
)

var testUser = model.User{
	Id:    idTest,
	Name:  nameTest,
	Email: emailTest,
}

func TestServer_FindUser(t *testing.T) {
	t.Run("No error return", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		serverTest := New(mockService)

		mockService.EXPECT().
			FindUser(gomock.Any(), idTest).
			Return(testUser, nil).
			Times(1)

		serverReturn := httptest.NewServer(serverTest)
		serverURL := fmt.Sprintf("%s/user/%s", serverReturn.URL, idTest)

		r, _ := http.NewRequest(http.MethodGet, serverURL, nil)
		w := httptest.NewRecorder()

		serverTest.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, string(w.Body.Bytes()), jsonOutput)
	})

	t.Run("Error GetUser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		mockService.EXPECT().
			FindUser(gomock.Any(), idTest).
			Return(model.User{}, fmt.Errorf("error")).
			Times(1)

		serverTest := New(mockService)

		serverReturn := httptest.NewServer(serverTest)
		serverURL := fmt.Sprintf("%s/user/%s", serverReturn.URL, idTest)

		r, _ := http.NewRequest(http.MethodGet, serverURL, nil)
		w := httptest.NewRecorder()

		serverTest.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestServer_CreateUser(t *testing.T) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(testUser)
	if err != nil {
		require.NoError(t, err)
	}

	t.Run("Create user successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		serverTest := New(mockService)

		mockService.EXPECT().
			CreateUser(gomock.Any(), testUser).
			Return(testUser, nil).
			Times(1)

		serverReturn := httptest.NewServer(serverTest)
		serverURL := fmt.Sprintf("%s/users", serverReturn.URL)

		reader := bytes.NewReader(buf.Bytes())

		r, _ := http.NewRequest(http.MethodPost, serverURL, reader)
		w := httptest.NewRecorder()

		serverTest.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Contains(t, string(w.Body.Bytes()), string(buf.Bytes()))
	})

	t.Run("Error creating user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		mockService.EXPECT().
			CreateUser(gomock.Any(), testUser).
			Return(model.User{}, fmt.Errorf("error")).
			Times(1)

		serverTest := New(mockService)

		serverReturn := httptest.NewServer(serverTest)
		serverURL := fmt.Sprintf("%s/users", serverReturn.URL)

		reader := bytes.NewReader(buf.Bytes())

		r, _ := http.NewRequest(http.MethodPost, serverURL, reader)
		w := httptest.NewRecorder()

		serverTest.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestServer_UpdateUser(t *testing.T) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(testUser)
	if err != nil {
		require.NoError(t, err)
	}

	t.Run("Update user successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		serverTest := New(mockService)

		mockService.EXPECT().
			UpdateUser(gomock.Any(), testUser).
			Return(testUser, nil).
			Times(1)

		serverReturn := httptest.NewServer(serverTest)
		serverURL := fmt.Sprintf("%s/user", serverReturn.URL)

		reader := bytes.NewReader(buf.Bytes())

		r, _ := http.NewRequest(http.MethodPut, serverURL, reader)
		w := httptest.NewRecorder()

		serverTest.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Contains(t, string(w.Body.Bytes()), string(buf.Bytes()))
	})

	t.Run("Error updating user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		mockService.EXPECT().
			UpdateUser(gomock.Any(), testUser).
			Return(model.User{}, fmt.Errorf("error")).
			Times(1)

		serverTest := New(mockService)

		serverReturn := httptest.NewServer(serverTest)
		serverURL := fmt.Sprintf("%s/user", serverReturn.URL)

		reader := bytes.NewReader(buf.Bytes())

		r, _ := http.NewRequest(http.MethodPut, serverURL, reader)
		w := httptest.NewRecorder()

		serverTest.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestServer_DeleteUser(t *testing.T) {
	t.Run("Delete user successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		serverTest := New(mockService)

		mockService.EXPECT().
			DeleteUser(gomock.Any(), idTest).
			Return(nil).
			Times(1)

		serverReturn := httptest.NewServer(serverTest)
		serverURL := fmt.Sprintf("%s/user/%s", serverReturn.URL, idTest)

		r, _ := http.NewRequest(http.MethodDelete, serverURL, nil)
		w := httptest.NewRecorder()

		serverTest.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Contains(t, string(w.Body.Bytes()), userDeletedMessage)
	})

	t.Run("Error deleting user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		mockService.EXPECT().
			DeleteUser(gomock.Any(), idTest).
			Return(fmt.Errorf("error")).
			Times(1)

		serverTest := New(mockService)

		serverReturn := httptest.NewServer(serverTest)
		serverURL := fmt.Sprintf("%s/user/%s", serverReturn.URL, idTest)

		r, _ := http.NewRequest(http.MethodDelete, serverURL, nil)
		w := httptest.NewRecorder()

		serverTest.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
