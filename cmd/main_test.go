package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	delivery "microblog-api/auth/delivery/http"
	"microblog-api/auth/repositories"
	"microblog-api/auth/services"
	"microblog-api/profile/repositories/mock"
	services2 "microblog-api/profile/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
	assert.NoError(t, err)
	userRepo, err := repositories.NewPostgresRepository(db)
	profileRepo := &mock.ProfileRepositoryMock{}
	profileService := services2.NewProfileService(profileRepo)
	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
	r := gin.Default()
	delivery.RegisterHTTPEndpoints(r, userService)
	creds := &delivery.UserCredentials{
		Username: "username",
		Password: "password",
	}
	body, err := json.Marshal(creds)
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestE2ESignin(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
	assert.NoError(t, err)
	userRepo, err := repositories.NewPostgresRepository(db)
	profileRepo := &mock.ProfileRepositoryMock{}
	profileService := services2.NewProfileService(profileRepo)
	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
	r := gin.Default()
	delivery.RegisterHTTPEndpoints(r, userService)
	creds := &delivery.UserCredentials{
		Username: "username",
		Password: "password",
	}
	body, err := json.Marshal(creds)
	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
