package http

import (
	"imguessr/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	UserHandler *UserHandler
}

type UserHandler struct {
	UserSvc domain.UserSvc
}

func NewHandler(userSvc domain.UserSvc) *Handler {
	return &Handler{
		UserHandler: &UserHandler{
			UserSvc: userSvc,
		},
	}
}

func (h *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Not valid UUID")
		return
	}

	user, err := h.UserSvc.Get(userId.String())
	if err != nil {
		c.JSON(http.StatusNotFound, "Can't find user with this id")
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c *gin.Context) {
	var newUser domain.User

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.UserSvc.Create(&newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newUser)
}