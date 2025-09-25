package post

import "microblog-api/models"

type Service interface {
	GetByUserId(userId string) []models.Post
	GetById(id string) (*models.Post, error)
	Delete(userId, id string) error
	Create(content, userId string) error
}
