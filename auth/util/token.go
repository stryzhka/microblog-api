package util

import (
	"github.com/golang-jwt/jwt/v5"
	"microblog-api/auth"
	"microblog-api/models"
	"time"
)

func GenerateToken(id, role, signingKey string) (string, error) {
	claims := auth.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: &models.User{
			Id:       id,
			Role:     role,
			Username: "",
			Password: "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signingKey))
}
