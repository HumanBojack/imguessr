package http

import (
	"fmt"
	"imguessr/pkg/domain"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthSvc domain.AuthSvc
	UserSvc domain.UserSvc
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the "Authorization" header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}
		tokenString := authHeader[len("Bearer "):]

		// Validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Set the user ID and isAdmin in the Gin context
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["id"].(string)
		c.Set("userID", userID)
		isAdmin := claims["isAdmin"].(bool)
		c.Set("isAdmin", isAdmin)

		c.Next()
	}
}

func authorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the admin variable from the Gin context
		isAdmin, ok := c.Get("isAdmin")
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "isAdmin not found in Gin context"})
			return
		}

		// Check if the user is an admin
		if !isAdmin.(bool) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}

func checkIsAdmin(c *gin.Context) bool {
	// Get the admin variable from the Gin context
	isAdmin, ok := c.Get("isAdmin")
	if !ok {
		return false
	}

	// Check if the user is an admin
	return isAdmin.(bool)
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
