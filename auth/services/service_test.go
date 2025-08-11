package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	testMock "microblog-api/auth/repositories/mock"
	"microblog-api/auth/util"
	"microblog-api/models"
	"testing"
	"time"
)

func TestUserFlow(t *testing.T) {
	repo := new(testMock.UserRepositoryMock)
	s := NewUserService(repo, "salt", "key", time.Hour)
	username := "auth"
	password := "password"
	user := &models.User{
		Username: username,
		Password: util.GeneratePasswordHash(password, s.passwordSalt),
		Role:     "auth",
	}
	repo.On("Create", mock.MatchedBy(func(u *models.User) bool {
		return u.Username == username && u.Password == util.GeneratePasswordHash(password, s.passwordSalt) &&
			u.Role == "auth" && len(u.Id) == 36
	})).Return(nil)
	err := s.Signup(username, password)
	assert.NoError(t, err)

	repo.On("Get", username, util.GeneratePasswordHash(password, s.passwordSalt)).Return(user, nil)
	token, err := s.Signin(username, util.GeneratePasswordHash(password, s.passwordSalt))
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
