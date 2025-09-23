package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	delivery "microblog-api/auth/delivery/http"
	"microblog-api/auth/repositories"
	"microblog-api/auth/services"
	"microblog-api/models"
	http2 "microblog-api/profile/delivery/http"
	repositories2 "microblog-api/profile/repositories"
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
	profileRepo, err := repositories2.NewPostgresRepository(db)
	profileService := services2.NewProfileService(profileRepo)
	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
	r := gin.Default()
	delivery.RegisterHTTPEndpoints(r, userService)
	creds := &delivery.UserCredentials{
		Username: "22",
		Password: "password",
	}
	body, err := json.Marshal(creds)
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignin(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
	assert.NoError(t, err)
	userRepo, err := repositories.NewPostgresRepository(db)
	profileRepo := &mock.ProfileRepositoryMock{}
	profileService := services2.NewProfileService(profileRepo)
	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
	r := gin.Default()

	delivery.RegisterHTTPEndpoints(r, userService)
	r.GET("/api/endpoint", delivery.NewAuthMiddleware(userService), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	creds := &delivery.UserCredentials{
		Username: "22",
		Password: "password",
	}
	body, err := json.Marshal(creds)
	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	body, _ = io.ReadAll(w.Body)
	type token struct {
		Token string `json:"token"`
	}
	tok := &token{}
	_ = json.Unmarshal(body, tok)
	req, _ = http.NewRequest("GET", "/api/endpoint", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	req, _ = http.NewRequest("GET", "/api/endpoint", nil)
	//fmt.Println(tok.Token)
	req.Header.Set("Authorization", "Bearer "+tok.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSigninGetProfile(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
	assert.NoError(t, err)
	userRepo, err := repositories.NewPostgresRepository(db)
	profileRepo, err := repositories2.NewPostgresRepository(db)
	profileService := services2.NewProfileService(profileRepo)
	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
	r := gin.Default()
	middleware := delivery.NewAuthMiddleware(userService)
	delivery.RegisterHTTPEndpoints(r, userService)
	api := r.Group("/api", middleware)
	http2.RegisterHTTPEndpoints(api, profileService)

	creds := &delivery.UserCredentials{
		Username: "pidoras",
		Password: "password",
	}
	body, err := json.Marshal(creds)
	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	body, _ = io.ReadAll(w.Body)
	type token struct {
		Token string `json:"token"`
	}
	tok := &token{}
	_ = json.Unmarshal(body, tok)
	req, _ = http.NewRequest("GET", "/api/profile/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	//req, _ = http.NewRequest("GET", "/api/profile/1", nil)
	////fmt.Println(tok.Token)
	//req.Header.Set("Authorization", "Bearer "+tok.Token)
	//w = httptest.NewRecorder()
	//r.ServeHTTP(w, req)
	//assert.Equal(t, http.StatusOK, w.Code)
	req, _ = http.NewRequest("GET", "/api/profile/1b4bb6-22ae-4b77-a418-1f4a633a2b35", nil)
	//fmt.Println(tok.Token)
	req.Header.Set("Authorization", "Bearer "+tok.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSigninUpdateProfile(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
	assert.NoError(t, err)
	userRepo, err := repositories.NewPostgresRepository(db)
	profileRepo, err := repositories2.NewPostgresRepository(db)
	profileService := services2.NewProfileService(profileRepo)
	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
	r := gin.Default()
	middleware := delivery.NewAuthMiddleware(userService)
	delivery.RegisterHTTPEndpoints(r, userService)
	api := r.Group("/api", middleware)
	http2.RegisterHTTPEndpoints(api, profileService)

	creds := &delivery.UserCredentials{
		Username: "pidoras",
		Password: "password",
	}
	body, err := json.Marshal(creds)
	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	body, _ = io.ReadAll(w.Body)
	type token struct {
		Token string `json:"token"`
	}
	tok := &token{}
	_ = json.Unmarshal(body, tok)
	req, _ = http.NewRequest("GET", "/api/profile/1", nil)
	//fmt.Println(tok.Token)
	req.Header.Set("Authorization", "Bearer "+tok.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//fmt.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	newProfile := &models.Profile{
		Id:     "",
		UserId: "",
		Name:   "dolboeb",
		Photo:  "",
		Status: "bely voron",
	}
	body, err = json.Marshal(newProfile)
	assert.NoError(t, err)
	req, _ = http.NewRequest("PUT", "/api/profile/", bytes.NewBuffer(body))
	//fmt.Println(tok.Token)
	req.Header.Set("Authorization", "Bearer "+tok.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}
