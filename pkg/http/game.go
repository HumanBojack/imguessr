package http

import (
	"imguessr/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	GameSvc domain.GameSvc
	UserSvc domain.UserSvc
}

func (h *GameHandler) Create(c *gin.Context) {
	// Bind the JSON body to the GameParameters struct
	var game domain.Game

	if err := c.BindJSON(&game); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get the user from the context
	userID, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Use the userID to get the user from the database
	user, err := h.UserSvc.GetUserByID(userID.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Add the user to the game
	game.User = user

	// Remove the user from the UsersIDs slice
	for i, userID := range game.UsersIDs {
		if userID == user.ID {
			game.UsersIDs = append(game.UsersIDs[:i], game.UsersIDs[i+1:]...)
		}
	}

	// Verify the game parameters
	err = h.UserSvc.VerifyUsersIDs(game.UsersIDs)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.GameSvc.VerifyFrequency(game.Frequency)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.GameSvc.VerifySteps(game.Steps)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.GameSvc.VerifyHiderType(game.HiderType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.GameSvc.VerifyImage(game.Image)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create the game
	err = h.GameSvc.CreateGame(&game)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, game)
}
