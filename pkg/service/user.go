package service

import (
	"fmt"
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

// Verify that the UserIDs slice is not empty and that all the IDs are valid
func (us userSvc) VerifyUsersIDs(usersIDs []string) error {
	// Check if the UsersIDs slice is empty
	if len(usersIDs) == 0 {
		return fmt.Errorf("no users provided")
	}

	// Verify that game.UsersIDs values are real users
	for _, userID := range usersIDs {
		_, err := us.GetUserByID(userID)
		if err != nil {
			return fmt.Errorf("user with ID '%v' does not seem exist : %v", userID, err.Error())
		}
	}

	return nil
}
