package services

import (
	"github.com/google/uuid"
	"microblog-api/models"
	"microblog-api/profile"
)

type ProfileService struct {
	repo profile.Repository
}

func NewProfileService(repo profile.Repository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) Create(userId string, name string, photo string) error {
	profile := &models.Profile{
		Id:     uuid.New().String(),
		UserId: userId,
		Name:   name,
		Status: "",
		Photo:  "",
	}
	return s.repo.Create(profile)
}

func (s *ProfileService) GetById(id string) (*models.Profile, error) {
	return s.repo.GetById(id)
}

func (s *ProfileService) GetAll() []models.Profile {
	return s.repo.GetAll()
}

func (s *ProfileService) Update(id, userId string, newProfile *models.Profile) error {
	return s.repo.Update(id, userId, newProfile)
}
