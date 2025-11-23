package post

import (
	"context"
	"microblog-api/models"
	"microblog-api/storage"
)

type Service interface {
	GetByUserId(userId string) []models.Post
	GetById(id string) (*models.Post, error)
	GetAll() []models.Post
	Delete(userId, id string) error
	Create(ctx context.Context, content, userId string, photoData storage.FileData) error
	LikePost(postId, userId string) error
	DislikePost(postId, userId string) error
	AddComment(ctx context.Context, postId, userId, content string, photoData storage.FileData) error
}
