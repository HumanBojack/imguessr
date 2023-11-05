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

func (h *GameHandler) GetAll(c *gin.Context) {
	// userType defines which games to return
	userType := c.Query("userType")

	// Check the userID query parameter based on the user role
	var userID string
	if ok, _ := c.Get("isAdmin"); ok == true {
		userID = c.Query("userID")
	} else {
		if id, ok := c.Get("userID"); ok {
			if strID, ok := id.(string); ok {
				userID = strID
			}
		}
	}

	// Get the games
	var games []*domain.Game
	var err error
	if userID == "" {
		games, err = h.GameSvc.GetAllGames()
	} else {
		games, err = h.GameSvc.GetAllGamesByUserID(userType, userID)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, games)
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
