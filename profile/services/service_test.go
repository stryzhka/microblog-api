package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"microblog-api/models"
	testMock "microblog-api/profile/repositories/mock"
	"testing"
)

func TestProfileFlow(t *testing.T) {
	repo := &testMock.ProfileRepositoryMock{}
	s := NewProfileService(repo)
	name := "hui"
	userId := "1"
	profile := &models.Profile{
		Id:     "",
		UserId: userId,
		Name:   name,
		Status: "",
		Photo:  "",
	}
	repo.On("Create", mock.MatchedBy(func(p *models.Profile) bool {
		return len(p.Id) == 36 && p.Name == name && p.UserId == userId
	})).Return(nil)
	err := s.Create(userId, name, "")
	assert.NoError(t, err)
	repo.On("GetById", userId).Return(profile, nil)
	_, err = s.GetById(userId)
	assert.NoError(t, err)
}
