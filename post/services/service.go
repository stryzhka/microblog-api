package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"microblog-api/auth"
	"microblog-api/models"
	"microblog-api/post"
	"microblog-api/storage"
	"strings"
	"time"
)

type PostService struct {
	repo    post.Repository
	storage storage.FileStorage
}

func NewPostService(repo post.Repository, storage storage.FileStorage) *PostService {
	return &PostService{
		repo:    repo,
		storage: storage,
	}
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

func (s *PostService) GetAll() []models.Post {
	return s.repo.GetAll()
}

func (s *PostService) Create(ctx context.Context, content, userId string, photoData storage.FileData) error {
	id, err := uuid.NewRandom()
	if strings.TrimSpace(content) == "" {
		return auth.ErrValidation
	}
	if err != nil {
		return err
	}
	photoPath := ""
	if photoData.File != nil {
		filename := fmt.Sprintf("posts/%d_%s", time.Now().UnixNano(), userId)
		err := s.storage.UploadFile(ctx, photoData.File, photoData.ContentType, filename, photoData.Size)
		photoPath, err = s.storage.GetFileURL(ctx, filename)
		if err != nil {
			return fmt.Errorf("failed to upload photo: %w", err)
		}
	}
	post := &models.Post{
		Id:          id.String(),
		ProfileId:   userId,
		Content:     content,
		DateCreated: time.Now().Format(time.RFC3339),
		Likes:       0,
		PicturePath: photoPath,
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
