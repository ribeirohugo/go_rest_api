package mongodb

import (
	"context"
	"fmt"
	"github.com/ribeirohugo/golang_startup/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	userCollection = "users"
)

func (db *Database) FindUser(ctx context.Context, id string) (model.User, error) {
	var user model.User

	collection := db.client.Database(db.database).Collection(userCollection)

	cursor, err := collection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		return model.User{}, err
	}

	if cursor.RemainingBatchLength() == 0 {
		return model.User{}, fmt.Errorf("user not found")
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
	collection := db.client.Database(db.database).Collection(userCollection)

	update := bson.M{"$set": user}

	_, err := collection.UpdateByID(ctx, user.Id, update)

	return err
}

func (db *Database) CreateUser(ctx context.Context, user model.User) (string, error) {
	collection := db.client.Database(db.database).Collection(userCollection)

	id := primitive.NewObjectID().String()

	user.Id = id

	_, err := collection.InsertOne(ctx, user)

	return id, err
}

func (db *Database) DeleteUser(ctx context.Context, id string) error {
	collection := db.client.Database(db.database).Collection(userCollection)

	_, err := collection.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}
