package http

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"microblog-api/models"
	"microblog-api/profile"
	"microblog-api/storage"
	"net/http"
)

type Handler struct {
	s profile.Service
}

func NewHandler(s profile.Service) *Handler {
	return &Handler{s: s}
}

type ProfileData struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Photo  string `json:"photo"`
}

// GetById godoc
// @Summary Get profile by profile id
// @Tags profile
// @Produce json
// @Param id path string true "profile id"
// @Success 200 {object} ProfileData
// @Failure 500
// @Router /api/profile/{id} [get]
func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")
	p, err := h.s.GetById(id)
	if errors.Is(err, profile.ErrProfileNotFound) {
		c.JSON(http.StatusOK, gin.H{})
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	profileData := &ProfileData{
		UserId: p.UserId,
		Name:   p.Name,
		Status: p.Status,
		Photo:  p.Photo,
	}
	c.JSON(http.StatusOK, profileData)
}

// GetAll godoc
// @Summary Get all profiles
// @Tags profile
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} ProfileData
// @Failure 401
// @Router /api/profile/ [get]
func (h *Handler) GetAll(c *gin.Context) {
	p := h.s.GetAll()
	var profileData []ProfileData
	if len(p) == 0 {
		c.JSON(http.StatusOK, []ProfileData{})
		return
	}
	for _, profile := range p {
		temp := &ProfileData{
			UserId: profile.UserId,
			Name:   profile.Name,
			Status: profile.Status,
			Photo:  profile.Photo,
		}
		profileData = append(profileData, *temp)
	}
	c.JSON(http.StatusOK, profileData)
}

// Update godoc
// @Summary Update profile
// @Tags profile
// @Security ApiKeyAuth
// @Accept mpfd
// @Produce json
// @Param name formData string true "profile name"
// @Param status formData string false "profile status"
// @Param photo formData file false "photo"
// @Success 200
// @Failure 400
// @Failure 401
// @Router /api/profile [put]
func (h *Handler) Update(c *gin.Context) {

	profileName := c.PostForm("name")
	profileStatus := c.PostForm("status")
	//profilePhoto := c.PostForm("")
	user := c.Value("user").(*models.User)
	creds := &models.Profile{
		Id:     "",
		UserId: "",
		Name:   profileName,
		Status: profileStatus,
		Photo:  "",
	}
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
	err = h.s.Update(c.Request.Context(), "", user.Id, creds, photoData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
