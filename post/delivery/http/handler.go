package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"microblog-api/models"
	"microblog-api/post"
	post2 "microblog-api/post"
	"net/http"
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
	postData := &PostData{}
	if err := c.BindJSON(postData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := h.s.Create(postData.Content, c.Value("user").(*models.User).Id)
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
	c.JSON(http.StatusOK, nil)
}
