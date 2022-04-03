package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

func (db *Database) FindUser(ctx context.Context, id string) (model.User, error) {
	row := db.client.QueryRowContext(ctx, `
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
		Id:    uid.String,
		Name:  name.String,
		Email: email.String,
	}

	return user, nil
}

func (db *Database) UpdateUser(ctx context.Context, user model.User) error {
	err := db.client.QueryRowContext(ctx, `
		UPDATE users 
		SET username = $2, email = $3
		WHERE id = $1
	`, user.Id, user.Name, user.Email).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) CreateUser(ctx context.Context, user model.User) (string, error) {
	res, err := db.client.ExecContext(ctx, `
		INSERT INTO users (username, email)
		VALUES ($1, $2)
	`, user.Name, user.Email)
	if err != nil {
		return "", fmt.Errorf("error creating user: %s", err.Error())
	}

	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("error returning last insert ID: %s", err.Error())
	}

	log.Println(lastInsertedId)

	return strconv.FormatInt(lastInsertedId, 10), nil
}

func (db *Database) DeleteUser(ctx context.Context, id string) error {
	err := db.client.QueryRowContext(ctx, `
		DELETE FROM users
		WHERE id = $1
	`, id).Err()
	if err != nil {
		return err
	}

	return nil
}
