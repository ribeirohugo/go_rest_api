package postgres

import (
	"context"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

func (r *Database) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	// TODO: Add database queries here

	return model.User{}, nil
}
