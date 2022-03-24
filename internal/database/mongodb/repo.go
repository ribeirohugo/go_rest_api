package mongodb

import (
	"context"

	"github.com/ribeirohugo/golang_startup/internal/model"

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

	_, err := collection.InsertOne(ctx, user)

	return id.String(), err
}

func (db *Database) DeleteUser(ctx context.Context, id string) error {
	idStr, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := db.client.Database(db.database).Collection(userCollection)

	_, err = collection.DeleteOne(ctx, bson.M{
		"_id": idStr,
	})

	return err
}
