package domain

type User struct {
	ID string `bson:"_id,omitempty" json:"_id"`
	UpdateUser
}

// TODO: forbid duplicate email
type UpdateUser struct {
	Name     string `bson:"name" binding:"required" json:"name" unique:"true"`
	Password string `bson:"password" binding:"required" json:"password"`
	IsAdmin  bool   `bson:"isAdmin" json:"isAdmin"`
}

type UserDB interface {
	GetAllUsers() ([]*User, error)
	GetUserByID(id string) (*User, error)
	GetUserByName(name string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(u *User) error
	DeleteUser(id string) error
}

type UserSvc interface {
	GetAllUsers() ([]*User, error)
	GetUserByID(id string) (*User, error)
	GetUserByName(name string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(u *User) error
	DeleteUser(id string) error
}
