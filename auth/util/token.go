package util

import (
	"github.com/golang-jwt/jwt/v5"
	"microblog-api/auth"
	"time"
)

func GenerateToken(id, role, signingKey string) (string, error) {
	claims := auth.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Id:   id,
		Role: role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signingKey))
}
