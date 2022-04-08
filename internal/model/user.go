package model

// User - domain model for a user
type User struct {
	ID    string `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string `bson:"name,omitempty" json:"name,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}
