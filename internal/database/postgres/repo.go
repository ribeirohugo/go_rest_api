package postgres

import (
	"context"
	"database/sql"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

// FindUser - Returns a user for a given ID or an error if anything fails
func (db *Database) FindUser(ctx context.Context, id string) (model.User, error) {
	row := db.sql.QueryRowContext(ctx, `
		SELECT id, username, email
		FROM users WHERE id = $1
		LIMIT 1
	`, id)

	var uid, name, email sql.NullString

	err := row.Scan(&uid, &name, &email)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		ID:    uid.String,
		Name:  name.String,
		Email: email.String,
	}

	return user, nil
}

// UpdateUser - Updates a user and returns an error if anything fails
func (db *Database) UpdateUser(ctx context.Context, user model.User) error {
	err := db.sql.QueryRowContext(ctx, `
		UPDATE users 
		SET username = $2, email = $3
		WHERE id = $1
	`, user.ID, user.Name, user.Email).Err()
	if err != nil {
		return err
	}

	return nil
}

// CreateUser - Creates a user and returns its ID or an error, if anything fails
func (db *Database) CreateUser(ctx context.Context, user model.User) (string, error) {
	lastInsertedID := ""

	err := db.sql.QueryRowContext(ctx, `
		INSERT INTO users (id, username, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`, user.ID, user.Name, user.Email).Scan(&lastInsertedID)
	if err != nil {
		return "", err
	}

	return lastInsertedID, nil
}

// DeleteUser - Deletes a User for a given ID and could return an error if anything fails
func (db *Database) DeleteUser(ctx context.Context, id string) error {
	err := db.sql.QueryRowContext(ctx, `
		DELETE FROM users
		WHERE id = $1
	`, id).Err()
	if err != nil {
		return err
	}

	return nil
}

// FindAllUsers - Returns all users for a given limit and offset
func (db *Database) FindAllUsers(ctx context.Context, offset int64, limit int64) ([]model.User, error) {
	// TODO
	panic("Not implemented yet")
}
