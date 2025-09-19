package http

import (
	"github.com/gin-gonic/gin"
	"microblog-api/auth"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	s auth.Service
}

func NewAuthMiddleware(s auth.Service) gin.HandlerFunc {
	return (&AuthMiddleware{s: s}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user, err := m.s.ParseToken(headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == auth.ErrInvalidToken {
			status = http.StatusUnauthorized
		}
		c.AbortWithStatus(status)
	}
	c.Set("user", user)
}
