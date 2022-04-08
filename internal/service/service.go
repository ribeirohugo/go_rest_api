//go:generate mockgen -package service -source=service.go -destination service_mock.go

package service

import (
	"context"
	"fmt"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

// Repository abstracts the data access layer.
type Repository interface {
	FindUser(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (string, error)
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, id string) error
}

// Service represents the user domain service layer.
type Service struct {
	repo Repository
}

// New instantiates a Service
func New(repo Repository) *Service {
	return &Service{repo: repo}
}

// FindUser - Returns a new user for a given ID. Returns the user or an error if anything fails.
func (s *Service) FindUser(ctx context.Context, id string) (model.User, error) {
	user, err := s.repo.FindUser(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("fail to get user: %v", err)
	}

	return user, nil
}

// CreateUser - Creates a new user. Returns the created user or an error if anything fails.
func (s *Service) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	userID, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return model.User{}, fmt.Errorf("fail creating user: %v", err)
	}

	createdUser, err := s.repo.FindUser(ctx, userID)
	if err != nil {
		return model.User{}, fmt.Errorf("fail finding created user: %v", err)
	}

	return createdUser, nil
}

// UpdateUser - Updates an existing user. It returns the updated user or returns an error if anything fails.
func (s *Service) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		return model.User{}, fmt.Errorf("fail updating user: %v", err)
	}

	return user, nil
}

// DeleteUser - Removes a user from the database. It returns an error if anything fails.
func (s *Service) DeleteUser(ctx context.Context, id string) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("fail deleting user: %v", err)
	}

	return nil
}
