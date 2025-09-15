package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"microblog-api/auth/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	r := gin.Default()
	s := &services.UserServiceMock{}
	RegisterHTTPEndpoints(r, s)
	creds := &UserCredentials{
		Username: "username",
		Password: "password",
	}
	body, err := json.Marshal(creds)
	assert.NoError(t, err)
	s.On("Signup", creds.Username, creds.Password).Return(nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestSignin(t *testing.T) {
	r := gin.Default()
	s := &services.UserServiceMock{}
	RegisterHTTPEndpoints(r, s)
	creds := &UserCredentials{
		Username: "username",
		Password: "password",
	}
	body, err := json.Marshal(creds)
	assert.NoError(t, err)
	s.On("Signin", creds.Username, creds.Password).Return("jwt", nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"token\":\"jwt\"}", w.Body.String())
}
