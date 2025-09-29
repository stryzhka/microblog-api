package services

import (
	"github.com/google/uuid"
	"microblog-api/models"
	"microblog-api/post"
	"time"
)

type PostService struct {
	repo post.Repository
}

func NewPostService(repo post.Repository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) GetById(id string) (*models.Post, error) {
	return s.repo.GetById(id)
}

func (s *PostService) Delete(userId, id string) error {
	return s.repo.Delete(userId, id)
}

func (s *PostService) GetByUserId(userId string) []models.Post {
	return s.repo.GetByUserId(userId)
}

func (s *PostService) Create(content, userId string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	post := &models.Post{
		Id:          id.String(),
		ProfileId:   userId,
		Content:     content,
		DateCreated: time.Now().Format(time.RFC3339),
		Likes:       0,
	}
	return s.repo.Create(post)
}

func (s *PostService) LikePost(postId, userId string) error {
	like := &models.Like{
		ProfileId: userId,
		PostId:    postId,
	}
	return s.repo.LikePost(like)
}

func (s *PostService) DislikePost(postId, userId string) error {
	like := &models.Like{
		ProfileId: userId,
		PostId:    postId,
	}
	return s.repo.DislikePost(like)
}
