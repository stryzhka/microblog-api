package profile

import "microblog-api/models"

type Service interface {
	Create(userId string, name string, photo string) error
	GetById(id string) (*models.Profile, error)
}
