package post

import "microblog-api/models"

type Repository interface {
	GetByUserId(userId string) []models.Post
	GetById(id string) (*models.Post, error)
	GetAll() []models.Post
	Delete(userId, id string) error
	Create(post *models.Post) error
	LikePost(like *models.Like) error
	DislikePost(like *models.Like) error
	AddComment(comment *models.Post, commentData *models.CommentData) error
}
