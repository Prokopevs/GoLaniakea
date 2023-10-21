package handler

import (
	"net/http"

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
	var u CreatePostReq
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

func (h *Handler) GetPosts(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	res, err := h.serv.GetPosts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetPostById(c *gin.Context) {
	id := c.Param("id")

	res, err := h.serv.GetPostById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
