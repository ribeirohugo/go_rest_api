package postgres

import (
	"context"
	"database/sql"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

func (db *Database) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	rows, err := db.sql.QueryContext(ctx, `
		SELECT id, name, email FROM users WHERE email = $1
		LIMIT 1
	`, email)
	if err != nil {
		return model.User{}, err
	}

	var uid, name, emailSQL sql.NullString

	for rows.Next() {
		err = rows.Scan(&uid, &name, &email)
		if err != nil {
			return model.User{}, err
		}
	}

	user := model.User{
		Id:    uid.String,
		Name:  name.String,
		Email: emailSQL.String,
	}

	return user, nil
}
