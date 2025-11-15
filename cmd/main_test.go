package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	delivery "microblog-api/auth/delivery/http"
	"microblog-api/auth/repositories"
	"microblog-api/auth/services"
	repositories2 "microblog-api/profile/repositories"
	services2 "microblog-api/profile/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=root dbname=app sslmode=disable")
	assert.NoError(t, err)
	userRepo := repositories.NewPostgresRepository(db)
	profileRepo := repositories2.NewPostgresRepository(db)
	profileService := services2.NewProfileService(profileRepo)
	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
	r := gin.Default()
	delivery.RegisterHTTPEndpoints(r, userService)
	creds := &delivery.UserCredentials{
		Username: "ffds",
		Password: "  f",
	}
	body, err := json.Marshal(creds)
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

//func TestSignin(t *testing.T) {
//	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//	assert.NoError(t, err)
//	userRepo, err := repositories.NewPostgresRepository(db)
//	profileRepo := &mock.ProfileRepositoryMock{}
//	profileService := services2.NewProfileService(profileRepo)
//	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
//	r := gin.Default()
//
//	delivery.RegisterHTTPEndpoints(r, userService)
//	r.GET("/api/endpoint", delivery.NewAuthMiddleware(userService), func(c *gin.Context) {
//		c.Status(http.StatusOK)
//	})
//	creds := &delivery.UserCredentials{
//		Username: "22",
//		Password: "password",
//	}
//	body, err := json.Marshal(creds)
//	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json")
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code)
//	body, _ = io.ReadAll(w.Body)
//	type token struct {
//		Token string `json:"token"`
//	}
//	tok := &token{}
//	_ = json.Unmarshal(body, tok)
//	req, _ = http.NewRequest("GET", "/api/endpoint", nil)
//	w = httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusUnauthorized, w.Code)
//
//	req, _ = http.NewRequest("GET", "/api/endpoint", nil)
//	//fmt.Println(tok.Token)
//	req.Header.Set("Authorization", "Bearer "+tok.Token)
//	w = httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code)
//}

//	func TestSigninGetProfile(t *testing.T) {
//		db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//		assert.NoError(t, err)
//		userRepo, err := repositories.NewPostgresRepository(db)
//		profileRepo, err := repositories2.NewPostgresRepository(db)
//		profileService := services2.NewProfileService(profileRepo)
//		userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
//		r := gin.Default()
//		middleware := delivery.NewAuthMiddleware(userService)
//		delivery.RegisterHTTPEndpoints(r, userService)
//		api := r.Group("/api", middleware)
//		http2.RegisterHTTPEndpoints(api, profileService)
//
//		creds := &delivery.UserCredentials{
//			Username: "666",
//			Password: "password",
//		}
//		body, err := json.Marshal(creds)
//		req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
//		req.Header.Set("Content-Type", "application/json")
//		w := httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		assert.Equal(t, http.StatusOK, w.Code)
//		body, _ = io.ReadAll(w.Body)
//		type token struct {
//			Token string `json:"token"`
//		}
//		tok := &token{}
//		_ = json.Unmarshal(body, tok)
//		req, _ = http.NewRequest("GET", "/api/profiles/1", nil)
//		w = httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		assert.Equal(t, http.StatusUnauthorized, w.Code)
//
//		//req, _ = http.NewRequest("GET", "/api/profile/1", nil)
//		////fmt.Println(tok.Token)
//		//req.Header.Set("Authorization", "Bearer "+tok.Token)
//		//w = httptest.NewRecorder()
//		//r.ServeHTTP(w, req)
//		//assert.Equal(t, http.StatusOK, w.Code)
//		req, _ = http.NewRequest("GET", "/api/profile/2c5c18eb-27b6-4ea1-8fbb-7ca93748cff9", nil)
//		//fmt.Println(tok.Token)
//		req.Header.Set("Authorization", "Bearer "+tok.Token)
//		w = httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		fmt.Println(w.Body)
//		assert.Equal(t, http.StatusOK, w.Code)
//	}
//
//	func TestSigninUpdateProfile(t *testing.T) {
//		db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//		assert.NoError(t, err)
//		userRepo, err := repositories.NewPostgresRepository(db)
//		profileRepo, err := repositories2.NewPostgresRepository(db)
//		profileService := services2.NewProfileService(profileRepo)
//		userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
//		r := gin.Default()
//		middleware := delivery.NewAuthMiddleware(userService)
//		delivery.RegisterHTTPEndpoints(r, userService)
//		api := r.Group("/api", middleware)
//		http2.RegisterHTTPEndpoints(api, profileService)
//
//		creds := &delivery.UserCredentials{
//			Username: "777",
//			Password: "password",
//		}
//		body, err := json.Marshal(creds)
//		req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
//		req.Header.Set("Content-Type", "application/json")
//		w := httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		assert.Equal(t, http.StatusOK, w.Code)
//		body, _ = io.ReadAll(w.Body)
//		type token struct {
//			Token string `json:"token"`
//		}
//		tok := &token{}
//		_ = json.Unmarshal(body, tok)
//		req, _ = http.NewRequest("GET", "/api/profiles/1", nil)
//		//fmt.Println(tok.Token)
//		req.Header.Set("Authorization", "Bearer "+tok.Token)
//		w = httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		//fmt.Println(w.Body)
//		assert.Equal(t, http.StatusOK, w.Code)
//		newProfile := &models.Profile{
//			Id:     "",
//			UserId: "",
//			Name:   "dolboeb",
//			Photo:  "",
//			Status: "bely voron",
//		}
//		body, err = json.Marshal(newProfile)
//		assert.NoError(t, err)
//		req, _ = http.NewRequest("PUT", "/api/profiles/", bytes.NewBuffer(body))
//		//fmt.Println(tok.Token)
//		req.Header.Set("Authorization", "Bearer "+tok.Token)
//		w = httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		fmt.Println(w.Body)
//		assert.Equal(t, http.StatusOK, w.Code)
//	}
//
//	func TestSigninGetAllProfiles(t *testing.T) {
//		db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//		assert.NoError(t, err)
//		userRepo, err := repositories.NewPostgresRepository(db)
//		profileRepo, err := repositories2.NewPostgresRepository(db)
//		profileService := services2.NewProfileService(profileRepo)
//		userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
//		r := gin.Default()
//		middleware := delivery.NewAuthMiddleware(userService)
//		delivery.RegisterHTTPEndpoints(r, userService)
//		api := r.Group("/api", middleware)
//		http2.RegisterHTTPEndpoints(api, profileService)
//
//		creds := &delivery.UserCredentials{
//			Username: "666",
//			Password: "password",
//		}
//		body, err := json.Marshal(creds)
//		req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
//		req.Header.Set("Content-Type", "application/json")
//		w := httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		assert.Equal(t, http.StatusOK, w.Code)
//		body, _ = io.ReadAll(w.Body)
//		type token struct {
//			Token string `json:"token"`
//		}
//		tok := &token{}
//		_ = json.Unmarshal(body, tok)
//		req, _ = http.NewRequest("GET", "/api/profile/", nil)
//		//fmt.Println(tok.Token)
//		req.Header.Set("Authorization", "Bearer "+tok.Token)
//		w = httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		fmt.Println(w.Body)
//		assert.Equal(t, http.StatusOK, w.Code)
//
// }
//func TestSigninCreatePost(t *testing.T) {
//	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//	assert.NoError(t, err)
//	userRepo := repositories.NewPostgresRepository(db)
//	profileRepo := repositories2.NewPostgresRepository(db)
//	profileService := services2.NewProfileService(profileRepo)
//	postRepo := repositories3.NewPostgresRepository(db)
//	postService := services3.NewPostService(postRepo)
//	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
//	r := gin.Default()
//	middleware := delivery.NewAuthMiddleware(userService)
//
//	delivery.RegisterHTTPEndpoints(r, userService)
//	api := r.Group("/api")
//	http3.RegisterHTTPEndpoints(api, postService, middleware)
//	http4.RegisterHTTPEndpoints(api, profileService, postService, middleware)
//	creds := &delivery.UserCredentials{
//		Username: "666",
//		Password: "password",
//	}
//	body, err := json.Marshal(creds)
//	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json")
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code)
//	body, _ = io.ReadAll(w.Body)
//	type token struct {
//		Token string `json:"token"`
//	}
//	tok := &token{}
//	_ = json.Unmarshal(body, tok)
//	postData := http3.PostData{}
//	body, _ = json.Marshal(postData)
//	req, _ = http.NewRequest("POST", "/api/post/", bytes.NewBuffer(body))
//	//fmt.Println(tok.Token)
//	req.Header.Set("Authorization", "Bearer "+tok.Token)
//	w = httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	fmt.Println(w.Body)
//	assert.Equal(t, http.StatusOK, w.Code)
//}
//
//func TestSigninGetPostByProfile(t *testing.T) {
//	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//	assert.NoError(t, err)
//	userRepo := repositories.NewPostgresRepository(db)
//	profileRepo := repositories2.NewPostgresRepository(db)
//	profileService := services2.NewProfileService(profileRepo)
//	postRepo := repositories3.NewPostgresRepository(db)
//	postService := services3.NewPostService(postRepo)
//	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
//	r := gin.Default()
//	middleware := delivery.NewAuthMiddleware(userService)
//
//	delivery.RegisterHTTPEndpoints(r, userService)
//	api := r.Group("/api")
//	http3.RegisterHTTPEndpoints(api, postService, middleware)
//
//	creds := &delivery.UserCredentials{
//		Username: "666",
//		Password: "password",
//	}
//	body, err := json.Marshal(creds)
//	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json")
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code)
//	body, _ = io.ReadAll(w.Body)
//	type token struct {
//		Token string `json:"token"`
//	}
//	tok := &token{}
//	_ = json.Unmarshal(body, tok)
//	postData := http3.PostData{}
//	body, _ = json.Marshal(postData)
//	req, _ = http.NewRequest("GET", "/api/profile/posts/", bytes.NewBuffer(body))
//	//fmt.Println(tok.Token)
//	req.Header.Set("Authorization", "Bearer "+tok.Token)
//	w = httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	fmt.Println(w.Body)
//	assert.Equal(t, http.StatusOK, w.Code)
//}

//func TestGetPost(t *testing.T) {
//	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//	assert.NoError(t, err)
//	postRepo := repositories3.NewPostgresRepository(db)
//	postService := services3.NewPostService(postRepo)
//	r := gin.Default()
//	api := r.Group("/api")
//	http3.RegisterHTTPEndpoints(api, postService, nil)
//	req, _ := http.NewRequest("GET", "/api/post/49195280-b824-4b1e-90f6-3fbcad2159cc", nil)
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	fmt.Println(w.Body)
//	assert.Equal(t, http.StatusOK, w.Code)
//}

//func TestGetByUserId(t *testing.T) {
//	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//	assert.NoError(t, err)
//	postRepo := repositories3.NewPostgresRepository(db)
//	postService := services3.NewPostService(postRepo)
//	r := gin.Default()
//	api := r.Group("/api")
//	http3.RegisterHTTPEndpoints(api, postService)
//
//	req, _ := http.NewRequest("GET", "/api/posts/profile/d6e1800b-2c2a-4793-9e15-156d69b3715f", nil)
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	fmt.Println(w.Body)
//	assert.Equal(t, http.StatusOK, w.Code)
//}

//func TestSigninDeletePost(t *testing.T) {
//	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//	assert.NoError(t, err)
//	userRepo, err := repositories.NewPostgresRepository(db)
//	profileRepo, err := repositories2.NewPostgresRepository(db)
//	profileService := services2.NewProfileService(profileRepo)
//	postRepo := repositories3.NewPostgresRepository(db)
//	postService := services3.NewPostService(postRepo)
//	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
//	r := gin.Default()
//	middleware := delivery.NewAuthMiddleware(userService)
//
//	delivery.RegisterHTTPEndpoints(r, userService)
//	api := r.Group("/api", middleware)
//	http3.RegisterHTTPEndpoints(api, postService)
//
//	creds := &delivery.UserCredentials{
//		Username: "666",
//		Password: "password",
//	}
//	body, err := json.Marshal(creds)
//	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json")
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code)
//	body, _ = io.ReadAll(w.Body)
//	type token struct {
//		Token string `json:"token"`
//	}
//	tok := &token{}
//	_ = json.Unmarshal(body, tok)
//	//postData := http3.PostData{Content: "dhufh89dfh89fdh98"}
//	//body, _ = json.Marshal(postData)
//	req, _ = http.NewRequest("DELETE", "/api/posts/profile/b54b84ef-dd66-42af-8acd-52b7802ff14c", nil)
//	//fmt.Println(tok.Token)
//	req.Header.Set("Authorization", "Bearer "+tok.Token)
//	w = httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	fmt.Println(w.Body)
//	assert.Equal(t, http.StatusOK, w.Code)
//}
//
//func TestSigninLikePost(t *testing.T) {
//	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//	assert.NoError(t, err)
//	userRepo, err := repositories.NewPostgresRepository(db)
//	profileRepo, err := repositories2.NewPostgresRepository(db)
//	profileService := services2.NewProfileService(profileRepo)
//	postRepo := repositories3.NewPostgresRepository(db)
//	postService := services3.NewPostService(postRepo)
//	userService := services.NewUserService(userRepo, profileService, "salt", "key", 10000)
//	r := gin.Default()
//	middleware := delivery.NewAuthMiddleware(userService)
//
//	delivery.RegisterHTTPEndpoints(r, userService)
//	api := r.Group("/api", middleware)
//	http3.RegisterHTTPEndpoints(api, postService)
//
//	creds := &delivery.UserCredentials{
//		Username: "666",
//		Password: "password",
//	}
//	body, err := json.Marshal(creds)
//	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json")
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code)
//	body, _ = io.ReadAll(w.Body)
//	type token struct {
//		Token string `json:"token"`
//	}
//	tok := &token{}
//	_ = json.Unmarshal(body, tok)
//	//postData := http3.PostData{Content: "dhufh89dfh89fdh98"}
//	//body, _ = json.Marshal(postData)
//	req, _ = http.NewRequest("DELETE", "/api/posts/4284cd73-9824-4ffe-93a7-af303fba8ca3", nil)
//	//fmt.Println(tok.Token)
//	req.Header.Set("Authorization", "Bearer "+tok.Token)
//	w = httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	fmt.Println(w.Body)
//	assert.Equal(t, http.StatusOK, w.Code)
//}

//func TestGetPost(t *testing.T) {
//	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")
//	assert.NoError(t, err)
//	postRepo := repositories3.NewPostgresRepository(db)
//	endpoint := "localhost:9000"
//	client, err := minio.New(endpoint, &minio.Options{
//		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
//		Secure: false,
//	})
//	if err != nil {
//		t.Fatalf("failed to create minio client: %v", err)
//	}
//	storage := minio_st.NewMinioStorage(client, "app")
//	postService := services3.NewPostService(postRepo, storage)
//	r := gin.Default()
//	api := r.Group("/api")
//	http3.RegisterHTTPEndpoints(api, postService, nil)
//
//	req, _ := http.NewRequest("POST", "/api/post/test", nil)
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	fmt.Println(w.Body)
//	assert.Equal(t, http.StatusOK, w.Code)
//}
