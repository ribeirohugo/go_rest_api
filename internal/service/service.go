//go:generate mockgen -package service -source=service.go -destination service_mock.go

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

func (s *Service) FindUser(ctx context.Context, id string) (model.User, error) {
	// TODO
	return model.User{}, nil
}

func (s *Service) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	// TODO
	return model.User{}, nil
}

func (s *Service) UpdateUser(ctx context.Context, user model.User) error {
	// TODO
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	// TODO
	return nil
}
