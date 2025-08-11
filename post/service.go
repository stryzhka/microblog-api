package post

import "microblog-api/models"

type Service interface {
	GetById(id string) (*models.Post, error)
	//GetByUserId(userId string) ([]models.Post, error)
	Create(content, userId string) error
}
