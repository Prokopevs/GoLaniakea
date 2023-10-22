package handler

import (
	"fmt"
	"github/Prokopevs/GoLaniakea/internal/model"
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
	category := c.Query("category")
	page := c.Query("page")
	limit := c.Query("limit")

	res, total, err := h.serv.GetPosts(category, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := PostsResponse {
		Posts: res,
		Total: total,
	}

	fmt.Println(total)
	c.JSON(http.StatusOK, response)
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

func (h *Handler) DeletePostById(c *gin.Context) {
	id := c.Param("id")

	res, err := h.serv.DeletePostById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdatePost(c *gin.Context) {
	var u model.Post
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.serv.UpdatePost(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
