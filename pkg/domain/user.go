package domain

type User struct {
	ID       string `bson:"_id,omitempty" json:"_id"`
	Name     string `bson:"name" binding:"required" json:"name"`
	Password string `bson:"password" binding:"required" json:"password"`
}

type UserDB interface {
	Get(id string) (*User, error)
	Create(user *User) error
}

type UserSvc interface {
	Get(id string) (*User, error)
	Create(user *User) error
}
