package http

import (
	"github.com/gin-gonic/gin"
	"microblog-api/auth"
)

func RegisterHTTPEndpoints(router *gin.Engine, s auth.Service) {
	h := NewHandler(s)
	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/signup", h.Signup)
		authEndpoints.POST("/signin", h.Signin)
	}
}
