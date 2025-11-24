package http

import (
	"github.com/gin-gonic/gin"
	"microblog-api/auth"
	"microblog-api/models"
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

// Signup godoc
// @Summary Sign up user
// @Tags auth
// @Accept json
// @Param creds body UserCredentials required "user credentials"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /auth/signup [post]
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
		//c.AbortWithStatus(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// Signin godoc
// @Summary Sign in user
// @Tags auth
// @Accept json
// @Produce json
// @Param creds body UserCredentials required "user credentials"
// @Success 200
// @Failure 400
// @Router /auth/signin [post]
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

// GetUserId godoc
// @Summary Get user id
// @Tags auth
// @Security ApiKeyAuth
// @Produce json
// @Success 200
// @Router /auth/me [get]
func (h *Handler) GetUserId(c *gin.Context) {
	userId := c.Value("user").(*models.User).Id
	c.JSON(200, gin.H{"id": userId})
}
