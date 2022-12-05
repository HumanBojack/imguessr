package service

import (
	"imguessr/pkg/domain"

	"github.com/google/uuid"
)

type userSvc struct {
	DB domain.UserDB
}

func NewUserSvc(db domain.UserDB) domain.UserSvc {
	return userSvc{
		DB: db,
	}
}

func (us userSvc) Get(id string) (*domain.User, error) {
	return us.DB.Get(id)
}

func (us userSvc) Create(u *domain.User) error {
	u.ID = uuid.New().String()
	return us.DB.Create(u)
}