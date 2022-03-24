package postgres

import (
	"context"
	"github.com/ribeirohugo/golang_startup/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	userCollection = "users"
)

func (db *Database) FindUser(ctx context.Context, id string) (model.User, error) {
	var user model.User

	idStr, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	collection := db.client.Database(db.database).Collection(userCollection)

	cursor, err := collection.Find(ctx, bson.M{"_id": idStr})
	if err != nil {
		return model.User{}, err
	}

	for cursor.Next(ctx) {
		err = cursor.Decode(&user)
		if err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

func (db *Database) UpdateUser(ctx context.Context, user model.User) error {
	return nil
}

func (db *Database) CreateUser(ctx context.Context, user model.User) (string, error) {
	collection := db.client.Database(db.database).Collection(userCollection)

	id := primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, bson.D{
		{"_id", id},
		{"username", user.Name},
		{"email", user.Email},
		{"updated", primitive.Timestamp{T: uint32(time.Now().Unix())}},
	})

	return id.String(), err
}

func (db *Database) DeleteUser(ctx context.Context, id string) error {
	return nil
}
