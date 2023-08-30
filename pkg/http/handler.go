package http

import (
	"imguessr/pkg/domain"
)

type Handler struct {
	UserHandler *UserHandler
}

func NewHandler(userSvc domain.UserSvc) *Handler {
	return &Handler{
		UserHandler: &UserHandler{
			UserSvc: userSvc,
		},
	}
}
