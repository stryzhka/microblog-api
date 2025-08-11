package post

import "microblog-api/models"

type Repository interface {
	//List() []models.Post
	Create(post *models.Post) error
	GetById(id string) (*models.Post, error)
	//Delete(id string) error
}
