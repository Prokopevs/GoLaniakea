package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HTTP) GetPosts(c *gin.Context) {
	category := c.Query("category")
	page := c.Query("page")
	limit := c.Query("limit")

	res, total, err := h.serv.GetPosts(category, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := PostsResponse{
		Posts: res,
		Total: total,
	}

	fmt.Println(total)
	c.JSON(http.StatusOK, response)
}

func (h *HTTP) GetPostById(c *gin.Context) {
	id := c.Param("id")

	res, err := h.serv.GetPostById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}