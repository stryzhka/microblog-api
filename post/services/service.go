package services

import (
	"github.com/google/uuid"
	"microblog-api/models"
	"microblog-api/post"
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

func (s *PostService) Create(content, userId string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	post := &models.Post{
		Id:      id.String(),
		UserId:  userId,
		Content: content,
	}
	return s.repo.Create(post)
}
