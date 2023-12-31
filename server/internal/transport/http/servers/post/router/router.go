package router

import (
	"github/Prokopevs/GoLaniakea/internal/transport/http/servers/post/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(postHandler *handler.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/post/create", postHandler.CreatePost)
	r.GET("/posts", postHandler.GetPosts)
	r.GET("/post/:id", postHandler.GetPostById)
	r.DELETE("/post/delete/:id", postHandler.DeletePostById)
	r.PATCH("/post/update", postHandler.UpdatePost)
}

func Start(addr string) error {
	return r.Run(addr)
}
