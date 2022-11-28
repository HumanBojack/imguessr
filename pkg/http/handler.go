package http

import (
	"imguessr/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserSvc domain.UserSvc
}

func NewUserHandler(userSvc domain.UserSvc) *Handler {
	return &Handler{
		UserSvc: userSvc,
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Get user !!!")
}