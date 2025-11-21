package http

import (
	"github.com/gin-gonic/gin"
	"microblog-api/auth"
)

func RegisterHTTPEndpoints(router *gin.Engine, s auth.Service, m gin.HandlerFunc) {
	h := NewHandler(s)
	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/signup", h.Signup)
		authEndpoints.POST("/signin", h.Signin)
		authEndpoints.GET("/me", m, h.GetUserId)
	}
}
