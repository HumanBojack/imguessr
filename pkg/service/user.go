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

func (us userSvc) GetAll() ([]*domain.User, error) {
	return us.DB.GetAll()
}

func (us userSvc) Get(id string) (*domain.User, error) {
	return us.DB.Get(id)
}

func (us userSvc) GetByName(name string) (*domain.User, error) {
	return us.DB.GetByName(name)
}

func (us userSvc) Create(u *domain.User) error {
	u.ID = uuid.New().String()
	return us.DB.Create(u)
}

func (us userSvc) Update(u *domain.User) error {
	// TODO: sanity check (email, password...)
	err := us.DB.Update(u)
	return err
}

func (us userSvc) Delete(id string) error {
	err := us.DB.Delete(id)
	return err
}
