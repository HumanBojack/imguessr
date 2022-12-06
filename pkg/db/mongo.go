package db

import (
	"context"
	"imguessr/pkg/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStore struct {
	Client         *mongo.Client
	UserCollection *mongo.Collection
}

func NewMongoStore() (domain.UserDB, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo"))
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		return nil, err
	}
	return mongoStore{
		Client:         client,
		UserCollection: client.Database("main").Collection("users"),
	}, nil
}

func (ms mongoStore) Get(id string) (*domain.User, error) {
	result := ms.UserCollection.FindOne(context.TODO(), bson.M{"_id": id})

	var user domain.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ms mongoStore) Create(u *domain.User) error {
	_, err := ms.UserCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	return nil
}

func (ms mongoStore) Update(u *domain.User) error {
	filter := bson.M{"_id": u.ID}

	_, err := ms.UserCollection.ReplaceOne(
		context.TODO(),
		filter,
		u,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ms mongoStore) Delete(id string) error {
	filter := bson.M{"_id": id}
	_, err := ms.UserCollection.DeleteOne(
		context.TODO(),
		filter,
	)
	return err
}