package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/ribeirohugo/golang_startup/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	idTest    = "00000000-0000-0000-0000-000000000000"
	emailTest = "email@domain"
	nameTest  = "Test"
)

var userTest = model.User{
	Id:    idTest,
	Name:  nameTest,
	Email: emailTest,
}

func TestService_FindUser(t *testing.T) {
	t.Run("Returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)

		service := New(repositoryMock)

		repositoryMock.EXPECT().
			FindUser(gomock.Any(), idTest).
			Return(model.User{}, fmt.Errorf("error")).
			Times(1)

		_, err := service.FindUser(context.Background(), idTest)
		assert.Error(t, err)
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)

		service := New(repositoryMock)

		repositoryMock.EXPECT().
			FindUser(gomock.Any(), idTest).
			Return(userTest, nil).
			Times(1)

		user, err := service.FindUser(context.Background(), idTest)
		assert.NoError(t, err)
		assert.Equal(t, userTest, user)
	})
}

func TestService_CreateUser(t *testing.T) {
	t.Run("Returns an error ", func(t *testing.T) {
		t.Run("in create user method.", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repositoryMock := NewMockRepository(ctrl)

			service := New(repositoryMock)

			gomock.InOrder(
				repositoryMock.EXPECT().
					CreateUser(gomock.Any(), userTest).
					Return(idTest, fmt.Errorf("error")).
					Times(1),
			)

			_, err := service.CreateUser(context.Background(), userTest)
			assert.Error(t, err)
		})

		t.Run("in find user method", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repositoryMock := NewMockRepository(ctrl)

			service := New(repositoryMock)

			gomock.InOrder(
				repositoryMock.EXPECT().
					CreateUser(gomock.Any(), userTest).
					Return(idTest, nil).
					Times(1),
				repositoryMock.EXPECT().
					FindUser(gomock.Any(), idTest).
					Return(model.User{}, fmt.Errorf("error")).
					Times(1),
			)

			_, err := service.CreateUser(context.Background(), userTest)
			assert.Error(t, err)
		})
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)

		service := New(repositoryMock)

		gomock.InOrder(
			repositoryMock.EXPECT().
				CreateUser(gomock.Any(), userTest).
				Return(idTest, nil).
				Times(1),
			repositoryMock.EXPECT().
				FindUser(gomock.Any(), idTest).
				Return(userTest, nil).
				Times(1),
		)

		user, err := service.CreateUser(context.Background(), userTest)
		assert.NoError(t, err)
		assert.Equal(t, userTest, user)
	})
}

func TestService_UpdateUser(t *testing.T) {
	t.Run("Returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)

		service := New(repositoryMock)

		repositoryMock.EXPECT().
			UpdateUser(gomock.Any(), userTest).
			Return(fmt.Errorf("error")).
			Times(1)

		_, err := service.UpdateUser(context.Background(), userTest)
		assert.Error(t, err)
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)

		service := New(repositoryMock)

		repositoryMock.EXPECT().
			UpdateUser(gomock.Any(), userTest).
			Return(nil).
			Times(1)

		user, err := service.UpdateUser(context.Background(), userTest)
		assert.NoError(t, err)
		assert.Equal(t, userTest, user)
	})
}

func TestService_DeleteUser(t *testing.T) {
	t.Run("Returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)

		service := New(repositoryMock)

		repositoryMock.EXPECT().
			DeleteUser(gomock.Any(), idTest).
			Return(fmt.Errorf("error")).
			Times(1)

		err := service.DeleteUser(context.Background(), idTest)
		assert.Error(t, err)
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)

		service := New(repositoryMock)

		repositoryMock.EXPECT().
			DeleteUser(gomock.Any(), idTest).
			Return(nil).
			Times(1)

		err := service.DeleteUser(context.Background(), idTest)
		assert.NoError(t, err)
	})
}
