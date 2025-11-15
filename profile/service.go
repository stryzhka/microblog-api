package profile

import (
	"context"
	"microblog-api/models"
	"microblog-api/storage"
)

type Service interface {
	Create(userId string, name string) error
	GetById(id string) (*models.Profile, error)
	GetAll() []models.Profile
	Update(ctx context.Context, id string, userId string, newProfile *models.Profile, photoData storage.FileData) error
}
