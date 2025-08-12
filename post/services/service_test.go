package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"microblog-api/models"
	testMock "microblog-api/post/repositories/mock"
	"testing"
)

func TestPostFlow(t *testing.T) {
	repo := testMock.PostRepositoryMock{}
	s := NewPostService(&repo)
	postId := uuid.New().String()
	post := &models.Post{
		Id:          "",
		ProfileId:   "",
		Likes:       0,
		Content:     "666",
		PicturePath: "",
		DateCreated: "",
	}
	repo.On("Create", mock.MatchedBy(func(p *models.Post) bool {
		return len(p.Id) == 36 && p.Content == post.Content
	})).Return(nil)
	repo.On("GetById", postId).Return(post, nil)
	err := s.Create("666", postId)
	assert.NoError(t, err)

	_, err = s.GetById(postId)
	assert.NoError(t, err)
}
