package service

import (
	"imguessr/pkg/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userSvc struct {
	DB domain.UserDB
}

func NewUserSvc(db domain.UserDB) domain.UserSvc {
	return userSvc{
		DB: db,
	}
}

func (us userSvc) Get(id primitive.ObjectID) (*domain.User, error) {
	return us.DB.Get(id)
}