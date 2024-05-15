package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *HTTP) setRoutes(r *gin.Engine) {
	api := r.Group("/api/v1") 
	{
		api.POST("/create", h.CreatePost)
		api.DELETE("/deletePost/:id", h.DeletePostById)
		api.PATCH("/updatePost", h.UpdatePost)

		api.GET("/getPosts", h.GetPosts)
		api.GET("/getPost/:id", h.GetPostById)
		api.GET("/getTotalCount", h.GetTotalCount)
		api.GET("/getInteresting", h.GetInteresting)
	}
}
