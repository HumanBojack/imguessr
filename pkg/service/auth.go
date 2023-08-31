package service

import (
	"imguessr/pkg/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type authSvc struct {
}

func NewAuthSvc() domain.AuthSvc {
	return authSvc{}
}

func (as authSvc) GenerateToken(id string, isAdmin bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["isAdmin"] = isAdmin
	// Set token expiration (24 hours)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}
