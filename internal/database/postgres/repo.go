package postgres

import (
	"context"
	"database/sql"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

func (db *Database) FindUser(ctx context.Context, id string) (model.User, error) {
	row := db.sql.QueryRowContext(ctx, `
		SELECT id, name, email FROM users WHERE id = $1
		LIMIT 1
	`, id)

	var uid, name, email sql.NullString

	err := row.Scan(&uid, &name, &id)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Id:    uid.String,
		Name:  name.String,
		Email: email.String,
	}

	return user, nil
}

func (db *Database) UpdateUser(ctx context.Context, user model.User) error {
	err := db.sql.QueryRowContext(ctx, `
		UPDATE users 
		SET username = $1, email = $2
		WHERE id = $5
	`, user.Id, user.Name, user.Email).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) CreateUser(ctx context.Context, user model.User) (string, error) {
	lastInsertedId := ""
	err := db.sql.QueryRowContext(ctx, `
		INSERT INTO users (id, name, email)
		VALUES ($1, $2, $3)
	`, user.Id, user.Name, user.Email).Scan(&lastInsertedId)
	if err != nil {
		return "", err
	}

	return lastInsertedId, nil
}

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
