package mock

import (
	"github.com/stretchr/testify/mock"
	"microblog-api/models"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (s *UserRepositoryMock) Create(user *models.User) error {
	args := s.Called(user)
	return args.Error(0)
}

func (s *UserRepositoryMock) Get(username, password string) (*models.User, error) {
	args := s.Called(username, password)
	return args.Get(0).(*models.User), args.Error(1)
}
