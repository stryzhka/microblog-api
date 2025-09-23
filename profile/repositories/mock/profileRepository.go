package mock

import (
	"github.com/stretchr/testify/mock"
	"microblog-api/models"
)

type ProfileRepositoryMock struct {
	mock.Mock
}

func (s *ProfileRepositoryMock) Create(profile *models.Profile) error {
	args := s.Called(profile)
	return args.Error(0)
}

func (s *ProfileRepositoryMock) GetById(id string) (*models.Profile, error) {
	args := s.Called(id)
	return args.Get(0).(*models.Profile), args.Error(1)
}

func (s *ProfileRepositoryMock) Update(id, userId string, newProfile *models.Profile) error {
	args := s.Called(id, userId, newProfile)
	return args.Error(0)
}
