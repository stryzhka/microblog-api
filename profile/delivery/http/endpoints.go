package http

import (
	"github.com/gin-gonic/gin"
	"microblog-api/post"
	"microblog-api/post/delivery/http"
	"microblog-api/profile"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, profileService profile.Service, postService post.Service, m gin.HandlerFunc) {
	h := NewHandler(profileService)
	hPost := http.NewHandler(postService)
	profileEndpoints := router.Group("/profile")
	{
		profileEndpoints.GET("/:id", m, h.GetById)
		profileEndpoints.PUT("/", m, h.Update)
		profileEndpoints.GET("/", m, h.GetAll)
		profileEndpoints.GET("/posts/:userId", m, hPost.GetByUserId)
	}
}
