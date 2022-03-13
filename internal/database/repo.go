package database

import (
	"context"
	"database/sql"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

// UserRepo is a user data access repository.
type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	// TODO: Add database queries here

	return model.User{}, nil
}
