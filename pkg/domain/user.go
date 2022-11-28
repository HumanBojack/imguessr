package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	Password string             `bson:"password,omitempty" json:"password"`
}

type UserDB interface {
	Get(id primitive.ObjectID) (*User, error)
}

type UserSvc interface {
	Get(id primitive.ObjectID) (*User, error)
}
