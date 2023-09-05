package db

import (
	"context"
	"imguessr/pkg/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (ms mongoStore) GetAllUsers() ([]*domain.User, error) {
	cur, err := ms.UserCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var userList = []*domain.User{}
	for cur.Next(context.TODO()) {
		var elem domain.User
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		userList = append(userList, &elem)
	}

	return userList, nil
}

func (ms mongoStore) GetUserByID(id string) (*domain.User, error) {
	result := ms.UserCollection.FindOne(context.TODO(), bson.M{"_id": id})

	var user domain.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ms mongoStore) GetUserByName(name string) (*domain.User, error) {
	result := ms.UserCollection.FindOne(context.TODO(), bson.M{
		"updateuser.name": name,
	})

	var user domain.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ms mongoStore) CreateUser(u *domain.User) error {
	_, err := ms.UserCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	return nil
}

func (ms mongoStore) UpdateUser(u *domain.User) error {
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

func (ms mongoStore) DeleteUser(id string) error {
	filter := bson.M{"_id": id}
	_, err := ms.UserCollection.DeleteOne(
		context.TODO(),
		filter,
	)
	return err
}
