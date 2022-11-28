package db

import (
	"context"
	"imguessr/pkg/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStore struct {
	Client *mongo.Client
}

func NewMongoStore() (domain.UserDB, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo"))
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		return nil, err
	}
	return mongoStore{Client: client}, nil
}

func (ms mongoStore) Get(id primitive.ObjectID) (*domain.User, error) {
	return &domain.User{
		Name: "Jean Ballon",
		Password: "balloche",
	}, nil
}

// func Seed(c *mongo.Collection, ctx context.Context) {
// 	user := User{
// 		Name:     "J(a)son",
// 		Password: "mdpdefou1234",
// 	}

// 	_, err := c.InsertOne(ctx, user)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	var users []User
// 	cur, _ := c.Find(ctx, bson.D{})
// 	cur.All(ctx, &users)

// 	fmt.Println(users)
// }