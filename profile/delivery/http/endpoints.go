package http

import (
	"github.com/gin-gonic/gin"

	"microblog-api/profile"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, s profile.Service) {
	h := NewHandler(s)
	authEndpoints := router.Group("/profile")
	{
		authEndpoints.GET("/:id", h.GetById)
		//authEndpoints.POST("/signin", h.Signin)
	}
}
