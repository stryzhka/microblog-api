package http

import (
	"github.com/gin-gonic/gin"
	"microblog-api/profile"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, s profile.Service, m gin.HandlerFunc) {
	h := NewHandler(s)
	profileEndpoints := router.Group("/profile")
	{
		profileEndpoints.GET("/:id", m, h.GetById)
		profileEndpoints.PUT("/", m, h.Update)
		profileEndpoints.GET("/", m, h.GetAll)

	}
}
