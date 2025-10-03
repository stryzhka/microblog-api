package http

import (
	"github.com/gin-gonic/gin"
	"microblog-api/post"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, s post.Service, m gin.HandlerFunc) {
	h := NewHandler(s)
	postEndpoints := router.Group("/post")
	{
		//profileEndpoints.GET("/:id", h.GetById)
		//profileEndpoints.PUT("/", h.Update)
		//postEndpoints.GET("/", h.GetAll)
		postEndpoints.POST("/", m, h.Create)
		postEndpoints.GET("/:id", h.GetById)
		//postEndpoints.DELETE("/:id", h.Delete)
		postEndpoints.POST("/:postId", m, h.LikePost)
		postEndpoints.DELETE("/:postId", m, h.DislikePost)
	}
}
