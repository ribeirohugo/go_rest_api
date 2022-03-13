package service

import (
	"context"
	"fmt"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

// UserRepo abstracts the data access layer.
type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

// Service represents the user domain service layer.
type Service struct {
	repo UserRepo
}

func NewService(repo UserRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, fmt.Errorf("fail to get user: %v", err)
	}

	return user, nil
}
