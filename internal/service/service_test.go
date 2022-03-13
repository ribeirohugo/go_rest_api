package service

import (
	"context"
	"fmt"
	"github.com/ribeirohugo/golang_startup/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/golang/mock/gomock"
)

const (
	emailTest = "email@domain.com"
)

func TestService_GetUserByEmail(t *testing.T) {
	userTest := model.User{
		Email: emailTest,
	}

	t.Run("Returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepoMock := NewMockUserRepo(ctrl)

		service := NewService(userRepoMock)

		userRepoMock.EXPECT().
			GetUserByEmail(gomock.Any(), emailTest).
			Return(model.User{}, fmt.Errorf("error")).
			Times(1)

		_, err := service.GetUserByEmail(context.Background(), emailTest)
		assert.Error(t, err)
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepoMock := NewMockUserRepo(ctrl)

		service := NewService(userRepoMock)

		userRepoMock.EXPECT().
			GetUserByEmail(gomock.Any(), emailTest).
			Return(userTest, nil).
			Times(1)

		user, err := service.GetUserByEmail(context.Background(), emailTest)
		assert.NoError(t, err)
		assert.Equal(t, userTest, user)
	})
}
