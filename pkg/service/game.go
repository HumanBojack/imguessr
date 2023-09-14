package service

import (
	"imguessr/pkg/domain"
)

type gameSvc struct {
	DB domain.GameDB
}

func NewGameSvc(db domain.GameDB) domain.GameSvc {
	return gameSvc{
		DB: db,
	}
}
