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

func (us userSvc) GetAllUsers() ([]*domain.User, error) {
	return us.DB.GetAllUsers()
}

func (us userSvc) GetUserByID(id string) (*domain.User, error) {
	return us.DB.GetUserByID(id)
}

func (us userSvc) GetUserByName(name string) (*domain.User, error) {
	return us.DB.GetUserByName(name)
}

func (us userSvc) CreateUser(u *domain.User) error {
	u.ID = uuid.New().String()
	return us.DB.CreateUser(u)
}

func (us userSvc) UpdateUser(u *domain.User) error {
	// TODO: sanity check (email, password...)
	err := us.DB.UpdateUser(u)
	return err
}

func (us userSvc) DeleteUser(id string) error {
	err := us.DB.DeleteUser(id)
	return err
}
