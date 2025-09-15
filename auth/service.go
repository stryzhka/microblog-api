package auth

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.RegisteredClaims
	Id   string `json: "id"`
	Role string `json: "role"`
}

type Service interface {
	Signup(username, password string) error
	Signin(username, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}
