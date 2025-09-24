package profile

import "microblog-api/models"

type Repository interface {
	Create(profile *models.Profile) error
	GetById(id string) (*models.Profile, error)
	GetAll() []models.Profile
	Update(id string, userId string, newProfile *models.Profile) error
}
