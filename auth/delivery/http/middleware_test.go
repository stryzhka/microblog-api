package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"microblog-api/auth"
	"microblog-api/auth/services"
	"microblog-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	r := gin.Default()
	s := &services.UserServiceMock{}
	r.POST("/api/endpoint", NewAuthMiddleware(s), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	w = httptest.NewRecorder()
	s.On("ParseToken", "").Return(&models.User{}, auth.ErrInvalidToken)
	req.Header.Set("Authorization", "Bearer ")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	w = httptest.NewRecorder()
	s.On("ParseToken", "token").Return(&models.User{}, nil)

	req.Header.Set("Authorization", "Bearer token")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
