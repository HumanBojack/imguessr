package http

import (
	"imguessr/pkg/domain"
)

type GameHandler struct {
	GameSvc domain.GameSvc
	UserSvc domain.UserSvc
}
