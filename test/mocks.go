package test

import (
	"fmt"
	"imguessr/pkg/domain"
)

// Mock the domain.UserDB interface
type mockDB struct {
}

var DB []*domain.User = []*domain.User{
	{
		ID: "39c71853-6206-4eef-9f5b-7a1a90830ccc",
		UpdateUser: domain.UpdateUser{
			Name:     "testLogin",
			Password: "$2y$10$Paa49WkbPM1mKL1Aoi30HuoDMPRxDOHjQZV0T4k5ItAKi5WPtintm",
			IsAdmin:  false,
		},
	},
	{
		ID: "ec4e2897-4ca4-4694-94d7-96db81ec223f",
		UpdateUser: domain.UpdateUser{
			Name:     "testAdmin",
			Password: "$2y$10$Paa49WkbPM1mKL1Aoi30HuoDMPRxDOHjQZV0T4k5ItAKi5WPtintm",
			IsAdmin:  true,
		},
	},
	{
		ID: "eee42492-e135-4c9a-89dc-923b9239b816",
		UpdateUser: domain.UpdateUser{
			Name:     "toDelete",
			Password: "$2y$10$Paa49WkbPM1mKL1Aoi30HuoDMPRxDOHjQZV0T4k5ItAKi5WPtintm",
			IsAdmin:  false,
		},
	},
}

// Mock the functions of the database interface
func (mdb mockDB) GetAll() ([]*domain.User, error) {
	return DB, nil
}

func (mdb mockDB) Get(id string) (*domain.User, error) {
	// Find the user with the given ID
	for _, user := range DB {
		if user.ID == id {
			return user, nil
		}
	}
	// Return the error if the user is not found
	return nil, fmt.Errorf("can't find user with id : %v", id)
}

func (mdb mockDB) GetByName(name string) (*domain.User, error) {
	// Find the user with the given name
	for _, user := range DB {
		if user.Name == name {
			return user, nil
		}
	}
	// Return the error if the user is not found
	return nil, fmt.Errorf("can't find user with name : %v", name)
}

func (mdb mockDB) Create(u *domain.User) error {
	// Add the user to the database
	DB = append(DB, u)

	return nil
}

func (mdb mockDB) Update(u *domain.User) error {
	// Find the user with the given ID
	for i, user := range DB {
		if user.ID == u.ID {
			// Update the user
			DB[i] = u
			return nil
		}
	}
	// Return the error if the user is not found
	return fmt.Errorf("can't find user with id : %v", u.ID)
}

func (mdb mockDB) Delete(id string) error {
	// Find the user with the given ID
	for i, user := range DB {
		if user.ID == id {
			// Delete the user
			DB = append(DB[:i], DB[i+1:]...)
			return nil
		}
	}
	// Return the error if the user is not found
	return fmt.Errorf("can't find user with id : %v", id)
}
