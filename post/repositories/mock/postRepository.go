package mock

import (
	"github.com/stretchr/testify/mock"
	"microblog-api/models"
)

type PostRepositoryMock struct {
	mock.Mock
}

func (s *PostRepositoryMock) Create(post *models.Post) error {
	args := s.Called(post)
	return args.Error(0)
}

func (s *PostRepositoryMock) GetById(id string) (*models.Post, error) {
	args := s.Called(id)
	return args.Get(0).(*models.Post), args.Error(1)
}
