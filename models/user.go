package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	Password string             `bson:"password,omitempty" json:"password"`
}

func Seed(c *mongo.Collection, ctx context.Context) {
	user := User{
		Name:     "J(a)son",
		Password: "mdpdefou1234",
	}

	_, err := c.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []User
	cur, _ := c.Find(ctx, bson.D{})
	cur.All(ctx, &users)

	fmt.Println(users)
}
