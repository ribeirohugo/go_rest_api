package postgres

import (
	"context"
	"database/sql"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

func (db *Database) FindUser(ctx context.Context, id string) (model.User, error) {
	rows, err := db.sql.QueryContext(ctx, `
		SELECT id, name, email FROM users WHERE id = $1
		LIMIT 1
	`, id)
	if err != nil {
		return model.User{}, err
	}

	var uid, name, email sql.NullString

	for rows.Next() {
		err = rows.Scan(&uid, &name, &id)
		if err != nil {
			return model.User{}, err
		}
	}

	user := model.User{
		Id:    uid.String,
		Name:  name.String,
		Email: email.String,
	}

	return user, nil
}
