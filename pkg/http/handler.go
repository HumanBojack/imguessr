package http

import (
	"imguessr/pkg/domain"
)

type Handler struct {
	UserHandler *UserHandler
	AuthHandler *AuthHandler
}

func NewHandler(userSvc domain.UserSvc, authSvc domain.AuthSvc) *Handler {
	return &Handler{
		UserHandler: &UserHandler{
			UserSvc: userSvc,
		},
		AuthHandler: &AuthHandler{
			UserSvc: userSvc,
			AuthSvc: authSvc,
		},
	}
}
