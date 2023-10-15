package handler

import (
	"net/http"

	services "github/Prokopevs/GoLaniakea/internal/services/post"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	serv PostService
}

func NewHandler(s PostService) *Handler {
	return &Handler{
		serv: s,
	}
}

func (h *Handler) CreatePost(c *gin.Context) {
	var u services.CreatePostReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.serv.CreatePost(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
