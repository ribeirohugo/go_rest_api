package model

type User struct {
	Id    string `json,bson:"id"`
	Name  string `json,bson:"name"`
	Email string `json,bson:"email"`
}
