package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ribeirohugo/golang_startup/internal/model"
	"github.com/stretchr/testify/assert"
)

const (
	emailTest = "email@domain"
	idTest    = "00000000-0000-0000-0000-000000000000"
	nameTest  = "Test"

	jsonOutput = "{\"Id\":\"00000000-0000-0000-0000-000000000000\",\"Name\":\"Test\",\"Email\":\"email@domain\"}\n"
)

var testUser = model.User{
	Id:    idTest,
	Name:  nameTest,
	Email: emailTest,
}

func TestServer_GetSingleUserByEmail(t *testing.T) {
	t.Run("No error return", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockUserService(ctrl)

		serverTest := New(mockService)

		mockService.EXPECT().
			GetUserByEmail(gomock.Any(), emailTest).
			Return(testUser, nil).
			Times(1)

		serverReturn := httptest.NewServer(serverTest)
		serverURL := serverReturn.URL + "/users/email?email=email@domain"

		r, _ := http.NewRequest("GET", serverURL, nil)
		w := httptest.NewRecorder()

		query := r.URL.Query()
		query.Add("email", emailTest)

		serverTest.GetSingleUserByEmail(w, r)

		assert.Equal(t, w.Code, http.StatusOK)

		assert.Equal(t, string(w.Body.Bytes()), jsonOutput)
	})

	t.Run("Error ", func(t *testing.T) {
		t.Run("no email query string found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := NewMockUserService(ctrl)

			serverTest := New(mockService)

			serverReturn := httptest.NewServer(serverTest)
			serverURL := serverReturn.URL + "/users/email"

			r, _ := http.NewRequest("GET", serverURL, nil)
			w := httptest.NewRecorder()

			serverTest.GetSingleUserByEmail(w, r)

			assert.Equal(t, w.Code, http.StatusBadRequest)
		})

		t.Run("GetUserByEmail", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := NewMockUserService(ctrl)

			mockService.EXPECT().
				GetUserByEmail(gomock.Any(), emailTest).
				Return(model.User{}, fmt.Errorf("error")).
				Times(1)

			serverTest := New(mockService)

			serverReturn := httptest.NewServer(serverTest)
			serverURL := serverReturn.URL + "/users/email?email=email@domain"

			r, _ := http.NewRequest("GET", serverURL, nil)
			w := httptest.NewRecorder()

			serverTest.GetSingleUserByEmail(w, r)

			assert.Equal(t, w.Code, http.StatusInternalServerError)
		})
	})
}
