package services

import (
	"github.com/stretchr/testify/mock"
	"microblog-api/models"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) Signup(username, password string) error {
	args := m.Called(username, password)
	return args.Error(0)
}

func (m *UserServiceMock) Signin(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.Get(0).(string), args.Error(1)
}

func (m *UserServiceMock) ParseToken(accessToken string) (*models.User, error) {
	args := m.Called(accessToken)

	return args.Get(0).(*models.User), args.Error(1)
}
