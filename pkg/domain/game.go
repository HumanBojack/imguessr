package domain

type Game struct {
	ID       string `bson:"_id,omitempty" json:"_id"`
	User     *User  `bson:"user" json:"user"`
	DateTime int64  `bson:"date_time" json:"date_time"`
	GameParameters
}

type GameParameters struct {
	UsersIDs  []string `bson:"users_ids" binding:"required" json:"users_ids"`
	Frequency int      `bson:"frequency" binding:"required" json:"frequency"`
	Steps     int      `bson:"steps" binding:"required" json:"steps"`
	HiderType string   `bson:"hider_type" binding:"required" json:"hider_type"`
	Image     []byte   `bson:"image" binding:"required" json:"image"`
}

type GameDB interface {
	CreateGame(game *Game) error
}

type GameSvc interface {
	CreateGame(game *Game) error
}
