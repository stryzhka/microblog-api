package http

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"microblog-api/models"
	"microblog-api/post"
	post2 "microblog-api/post"
	"microblog-api/storage"
	"mime/multipart"
	"net/http"
	"strings"
)

type Handler struct {
	s post.Service
}

func NewHandler(s post.Service) *Handler {
	return &Handler{s: s}
}

type PostData struct {
	Content string `json:"content"`
}

type CreatePostRequest struct {
	Content string                `form:"content" binding:"required"`
	Photo   *multipart.FileHeader `form:"photo"`
}

// Create godoc
// @Summary Create post
// @Tags post
// @Security ApiKeyAuth
// @Accept mpfd
// @Produce json
// @Param content formData string true "post content"
// @Param photo formData file false "photo"
// @Success 200
// @Failure 400
// @Failure 401
// @Router /api/post [post]
func (h *Handler) Create(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 20<<20)

	var req CreatePostRequest
	if err := c.ShouldBind(&req); err != nil {
		if strings.Contains(err.Error(), "http: request body too large") {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file too large"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form data"})
		return
	}
	postData := &PostData{Content: req.Content}
	photoData := storage.FileData{
		File:        nil,
		Size:        0,
		ContentType: "",
	}
	photo := req.Photo
	if photo != nil {
		file, err := photo.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()
		photoData.ContentType = photo.Header.Get("Content-Type")
		fmt.Println(photoData.ContentType)
		photoData.File = bytes.NewReader(fileBytes)
		photoData.Size = int64(len(fileBytes))
		if !strings.Contains(photoData.ContentType, "image") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file must be image"})
			return
		}

	}
	err := h.s.Create(c.Request.Context(), postData.Content, c.Value("user").(*models.User).Id, photoData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// GetById godoc
// @Summary Get post by post id
// @Tags post
// @Produce json
// @Param id path string true "post id"
// @Success 200 {object} models.Post
// @Failure 500
// @Router /api/post/{id} [get]
func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")
	post, err := h.s.GetById(id)
	if errors.Is(err, post2.ErrPostNotFound) {
		c.JSON(http.StatusOK, gin.H{})
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, post)
}

// GetById godoc
// @Summary Get posts by profile id
// @Tags profile
// @Security ApiKeyAuth
// @Produce json
// @Param userId path string true "user id"
// @Success 200 {array} models.Post
// @Failure 401
// @Router /api/profile/posts/{userId} [get]
func (h *Handler) GetByUserId(c *gin.Context) {
	userId := c.Param("userId")
	var posts []models.Post
	posts = h.s.GetByUserId(userId)
	if len(posts) == 0 {
		c.JSON(http.StatusOK, nil)
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetAll godoc
// @Summary Get all posts
// @Tags post
// @Produce json
// @Success 200 {array} models.Post
// @Router /api/post/ [get]
func (h *Handler) GetAll(c *gin.Context) {
	var posts []models.Post
	posts = h.s.GetAll()
	if len(posts) == 0 {
		c.JSON(http.StatusOK, nil)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *Handler) Delete(c *gin.Context) {
	userId := c.Value("user").(*models.User).Id
	id := c.Param("id")
	err := h.s.Delete(userId, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// LikePost godoc
// @Summary Like post by post id
// @Tags post
// @Security ApiKeyAuth
// @Produce json
// @Param postId path string true "post id"
// @Success 200
// @Failure 400
// @Failure 401
// @Router /api/post/{postId} [post]
func (h *Handler) LikePost(c *gin.Context) {
	userId := c.Value("user").(*models.User).Id
	postId := c.Param("postId")
	err := h.s.LikePost(postId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// DislikePost godoc
// @Summary Dislike post by post id
// @Tags post
// @Security ApiKeyAuth
// @Produce json
// @Param postId path string true "post id"
// @Success 200
// @Failure 400
// @Failure 401
// @Router /api/post/{postId} [delete]
func (h *Handler) DislikePost(c *gin.Context) {
	userId := c.Value("user").(*models.User).Id
	postId := c.Param("postId")
	err := h.s.DislikePost(postId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// AddComment godoc
// @Summary Add comment to post
// @Tags post
// @Security ApiKeyAuth
// @Accept mpfd
// @Produce json
// @Param postId path string true "post id"
// @Param content formData string true "comment content"
// @Param photo formData file false "photo"
// @Success 200
// @Failure 400
// @Failure 401
// @Router /api/post/comment/{postId} [post]
func (h *Handler) AddComment(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 20<<20)
	postId := c.Param("postId")
	var req CreatePostRequest
	if err := c.ShouldBind(&req); err != nil {
		if strings.Contains(err.Error(), "http: request body too large") {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file too large"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form data"})
		return
	}
	postData := &PostData{Content: req.Content}
	photoData := storage.FileData{
		File:        nil,
		Size:        0,
		ContentType: "",
	}
	photo := req.Photo
	if photo != nil {
		file, err := photo.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()
		photoData.ContentType = photo.Header.Get("Content-Type")
		fmt.Println(photoData.ContentType)
		photoData.File = bytes.NewReader(fileBytes)
		photoData.Size = int64(len(fileBytes))
		if !strings.Contains(photoData.ContentType, "image") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file must be image"})
			return
		}

	}
	err := h.s.AddComment(c.Request.Context(), postId, c.Value("user").(*models.User).Id, postData.Content, photoData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
