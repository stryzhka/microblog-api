package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"microblog-api/models"
	"microblog-api/profile/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	testUser := &models.User{
		Id:       "1",
		Role:     "auth",
		Username: "",
		Password: "",
	}
	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set("user", testUser)
	})
	s := &services.ProfileServiceMock{}
	RegisterHTTPEndpoints(group, s)
	s.On("GetById", "1").Return(&models.Profile{}, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/profile/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreate(t *testing.T) {
	testUser := &models.User{
		Id:       "1",
		Role:     "auth",
		Username: "",
		Password: "",
	}
	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set("user", testUser)
	})
	s := &services.ProfileServiceMock{}
	RegisterHTTPEndpoints(group, s)
	s.On("GetById", "1").Return(&models.Profile{}, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/profile/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
