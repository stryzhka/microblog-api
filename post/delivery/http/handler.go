package http

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"microblog-api/models"
	"microblog-api/post"
	post2 "microblog-api/post"
	"microblog-api/storage"
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

func (h *Handler) Create(c *gin.Context) {
	content := c.PostForm("content")
	if content == "" || strings.TrimSpace(content) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field 'content' is required"})
		return
	}
	postData := &PostData{Content: content}
	photoData := storage.FileData{
		File:        nil,
		Size:        0,
		ContentType: "",
	}
	photo, err := c.FormFile("photo")
	if err == nil {
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
		photoData.File = bytes.NewReader(fileBytes)
		photoData.Size = int64(len(fileBytes))
		photoData.ContentType = photo.Header.Get("Content-Type")
	}
	err = h.s.Create(c.Request.Context(), postData.Content, c.Value("user").(*models.User).Id, photoData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

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
