package http

import (
	"imguessr/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthSvc domain.AuthSvc
	UserSvc domain.UserSvc
}

func (h *AuthHandler) Login(c *gin.Context) {
	// Get login and password from the body
	var login domain.UpdateUser

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if the user exists
	user, err := h.UserSvc.GetByName(login.Name)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Generate a JWT token
	token, err := h.AuthSvc.GenerateToken(user.ID, user.IsAdmin)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
