package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"microblog-api/models"
	testMock "microblog-api/user/repositories/mock"
	"microblog-api/user/util"
	"testing"
	"time"
)

func TestUserFlow(t *testing.T) {
	repo := new(testMock.UserRepositoryMock)
	s := NewUserService(repo, "salt", "key", time.Hour)
	username := "user"
	password := "password"
	user := &models.User{
		Username: username,
		Password: util.GeneratePasswordHash(password, s.passwordSalt),
		Role:     "user",
	}
	repo.On("Create", mock.MatchedBy(func(u *models.User) bool {
		return u.Username == username && u.Password == util.GeneratePasswordHash(password, s.passwordSalt) &&
			u.Role == "user" && len(u.Id) == 36
	})).Return(nil)
	err := s.Signup(username, password)
	assert.NoError(t, err)

	repo.On("Get", username, util.GeneratePasswordHash(password, s.passwordSalt)).Return(user, nil)
	token, err := s.Signin(username, util.GeneratePasswordHash(password, s.passwordSalt))
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
