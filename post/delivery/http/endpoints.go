package http

import (
	"github.com/gin-gonic/gin"
	"microblog-api/post"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, s post.Service) {
	h := NewHandler(s)
	postEndpoints := router.Group("/post")
	{
		//profileEndpoints.GET("/:id", h.GetById)
		//profileEndpoints.PUT("/", h.Update)
		//profileEndpoints.GET("/", h.GetAll)
		postEndpoints.POST("/", h.Create)
		postEndpoints.GET("/:id", h.GetById)
		//postEndpoints.DELETE("/:id", h.Delete)
		postEndpoints.POST("/:postId", h.LikePost)
		postEndpoints.DELETE("/:postId", h.DislikePost)
	}
}
