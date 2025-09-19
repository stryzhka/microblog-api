package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"microblog-api/models"
)

type UserClaims struct {
	jwt.RegisteredClaims
	User *models.User `json:"user"`
}

type Service interface {
	Signup(username, password string) error
	Signin(username, password string) (string, error)
	ParseToken(accessToken string) (*models.User, error)
}
