package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"microblog-api/models"
	"microblog-api/profile"
	"net/http"
)

type Handler struct {
	s profile.Service
}

func NewHandler(s profile.Service) *Handler {
	return &Handler{s: s}
}

type ProfileData struct {
	//UserId string
	Name   string `json:"name"`
	Status string `json:"status"`
	Photo  string `json:"photo"`
}

func (h *Handler) GetById(c *gin.Context) {
	//
	id := c.Param("id")
	p, err := h.s.GetById(id)
	if errors.Is(err, profile.ErrProfileNotFound) {
		c.JSON(http.StatusOK, gin.H{})
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	profileData := &ProfileData{
		//UserId: p.UserId,
		Name:   p.Name,
		Status: p.Status,
		Photo:  p.Photo,
	}
	c.JSON(http.StatusOK, profileData)
}

func (h *Handler) Update(c *gin.Context) {
	creds := &models.Profile{}
	if err := c.BindJSON(creds); err != nil {
		fmt.Println("err", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user := c.Value("user").(*models.User)
	err := h.s.Update("", user.Id, creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
