package model

type User struct {
	Id    string `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string `bson:"name,omitempty" json:"name,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}
