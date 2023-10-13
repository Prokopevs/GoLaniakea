package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)

type Handler struct {
	service PostService 
}

func NewHandler(s PostService) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	// var u CreateUserReq
	// if err := c.ShouldBindJSON(&u); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// res, err := h.Service.CreateUser(c.Request.Context(), &u)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, res)
}