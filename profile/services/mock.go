package services

import (
	"github.com/stretchr/testify/mock"
	"microblog-api/models"
)

type ProfileServiceMock struct {
	mock.Mock
}

func (m *ProfileServiceMock) Create(userId string, name string, photo string) error {
	args := m.Called(userId, name, photo)
	return args.Error(0)
}

func (m *ProfileServiceMock) GetById(id string) (*models.Profile, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Profile), args.Error(1)
}
