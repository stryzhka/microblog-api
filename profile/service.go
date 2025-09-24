package profile

import "microblog-api/models"

type Service interface {
	Create(userId string, name string, photo string) error
	GetById(id string) (*models.Profile, error)
	GetAll() []models.Profile
	Update(id string, userId string, newProfile *models.Profile) error
}
