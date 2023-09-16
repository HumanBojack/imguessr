package service

import (
	"imguessr/pkg/domain"
	"time"

	"github.com/google/uuid"
)

type gameSvc struct {
	DB domain.GameDB
}

func NewGameSvc(db domain.GameDB) domain.GameSvc {
	return gameSvc{
		DB: db,
	}
}

func (gs gameSvc) CreateGame(g *domain.Game) error {
	g.ID = uuid.New().String()
	g.DateTime = time.Now().UTC().Unix()

	return gs.DB.CreateGame(g)
}
