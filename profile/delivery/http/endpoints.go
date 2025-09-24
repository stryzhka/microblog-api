package http

import (
	"github.com/gin-gonic/gin"

	"microblog-api/profile"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, s profile.Service) {
	h := NewHandler(s)
	profileEndpoints := router.Group("/profile")
	{
		profileEndpoints.GET("/:id", h.GetById)
		profileEndpoints.PUT("/", h.Update)
		//authEndpoints.POST("/signin", h.Signin)
	}
}
