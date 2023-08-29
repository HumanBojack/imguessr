package domain

type User struct {
	ID string `bson:"_id,omitempty" json:"_id"`
	UpdateUser
}

// TODO: forbid duplicate email
type UpdateUser struct {
	Name     string `bson:"name" binding:"required" json:"name" unique:"true"`
	Password string `bson:"password" binding:"required" json:"password"`
}

type UserDB interface {
	GetAll() ([]*User, error)
	Get(id string) (*User, error)
	Create(user *User) error
	Update(u *User) error
	Delete(id string) error
}

type UserSvc interface {
	GetAll() ([]*User, error)
	Get(id string) (*User, error)
	Create(user *User) error
	Update(u *User) error
	Delete(id string) error
}
