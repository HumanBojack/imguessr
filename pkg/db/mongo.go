package db

import (
	"context"
	"imguessr/pkg/domain"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStore struct {
	Client         *mongo.Client
	UserCollection *mongo.Collection
	GameCollection *mongo.Collection
}

func NewMongoStore() (domain.AppDB, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		return nil, err
	}
	ms := mongoStore{
		Client:         client,
		UserCollection: client.Database("main").Collection("users"),
		GameCollection: client.Database("main").Collection("games"),
	}

	// Create a unique index on the name field.
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"updateuser.name": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err = ms.UserCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return nil, err
	}

	return ms, nil
}
