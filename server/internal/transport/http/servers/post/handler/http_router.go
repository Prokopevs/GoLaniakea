package handler

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *HTTP) setRoutes(r *gin.Engine) {
	r.Use(CORSMiddleware())

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
