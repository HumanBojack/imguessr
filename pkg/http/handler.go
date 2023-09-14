package http

import (
	"imguessr/pkg/domain"
)

type Handler struct {
	UserHandler *UserHandler
	AuthHandler *AuthHandler
	GameHandler *GameHandler
}

func NewHandler(userSvc domain.UserSvc, authSvc domain.AuthSvc, gameSvc domain.GameSvc) *Handler {
	return &Handler{
		UserHandler: &UserHandler{
			UserSvc: userSvc,
		},
		AuthHandler: &AuthHandler{
			UserSvc: userSvc,
			AuthSvc: authSvc,
		},
		GameHandler: &GameHandler{
			GameSvc: gameSvc,
			UserSvc: userSvc,
		},
	}
}
