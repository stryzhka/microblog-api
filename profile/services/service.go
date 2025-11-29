package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"microblog-api/auth"
	"microblog-api/models"
	"microblog-api/profile"
	"microblog-api/storage"
	"strings"
	"time"
)

type ProfileService struct {
	repo    profile.Repository
	storage storage.FileStorage
}

func NewProfileService(repo profile.Repository, storage storage.FileStorage) *ProfileService {
	return &ProfileService{
		repo:    repo,
		storage: storage,
	}
}

func (s *ProfileService) Create(userId string, name string) error {

	profile := &models.Profile{
		Id:     uuid.New().String(),
		UserId: userId,
		Name:   name,
		Status: "",
	}
	return s.repo.Create(profile)
}

func (s *ProfileService) GetById(id string) (*models.Profile, error) {
	return s.repo.GetById(id)
}

func (s *ProfileService) GetAll() []models.Profile {
	return s.repo.GetAll()
}

func (s *ProfileService) Update(ctx context.Context, id, userId string, newProfile *models.Profile, photoData storage.FileData) error {
	if strings.TrimSpace(newProfile.Name) == "" {
		return auth.ErrValidation
	}
	if len(newProfile.Name) > 30 {
		return auth.ErrValidation
	}
	if len(newProfile.Status) > 30 {
		return auth.ErrValidation
	}
	photoPath := ""
	if photoData.File != nil {
		filename := fmt.Sprintf("profiles/%d_%s", time.Now().UnixNano(), userId)
		err := s.storage.UploadFile(ctx, photoData.File, photoData.ContentType, filename, photoData.Size)
		photoPath, err = s.storage.GetFileURL(ctx, filename)
		if err != nil {
			return fmt.Errorf("failed to upload photo: %w", err)
		}
	}

	newProfile.Photo = photoPath

	return s.repo.Update(id, userId, newProfile)
}
