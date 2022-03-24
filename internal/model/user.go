package model

type User struct {
	Id    string `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string `bson:"username,omitempty" json:"username,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}
