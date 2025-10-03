package http

import (
	"github.com/gin-gonic/gin"
	"microblog-api/auth"
	"net/http"
)

type Handler struct {
	s auth.Service
}

func NewHandler(s auth.Service) *Handler {
	return &Handler{s: s}
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Signup(c *gin.Context) {
	creds := &UserCredentials{}
	if err := c.BindJSON(creds); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.s.Signup(creds.Username, creds.Password)
	if err == auth.ErrUserAlreadyExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": auth.ErrUserAlreadyExists.Error()})
	} else if err == auth.ErrValidation {
		c.AbortWithStatus(http.StatusBadRequest)
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Signin(c *gin.Context) {
	creds := &UserCredentials{}
	if err := c.BindJSON(creds); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	token, err := h.s.Signin(creds.Username, creds.Password)
	if err == auth.ErrUserAlreadyExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": auth.ErrUserAlreadyExists.Error()})
	} else if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
