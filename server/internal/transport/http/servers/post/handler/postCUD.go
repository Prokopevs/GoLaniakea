package handler

import (
	"errors"
	"github/Prokopevs/GoLaniakea/internal/model"
	"github/Prokopevs/GoLaniakea/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostJSON struct {
	ImageUrl     string `json:"imageUrl"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Date         string `json:"date"`
	Category     string `json:"category"`
	CategoryName string `json:"categoryName"`
	LikeCount    string `json:"likeCount"`
	Text         string `json:"text"`
}

type response interface {
	writeJSON(*gin.Context)
}

func (h *HTTP) CreatePost(c *gin.Context) {
	resp := h.CreatePostResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) CreatePostResponse(c *gin.Context) response {
	p := &PostJSON{}
	if err := c.ShouldBindJSON(p); err != nil {
		return getBadRequestWithMsgResponse(codeBadBody, err.Error())
	}

	code, err := p.validate()
	if err != nil {
		return getBadRequestWithMsgResponse(code, err.Error())
	}

	post, err := convertPostJSONToPost(p)
	if err != nil {
		return getBadRequestWithMsgResponse(codeInvalidConvertion, err.Error())
	}

	result, err := h.service.CreatePost(c.Request.Context(), post)
	if err != nil {
		h.log.Errorw("Create post.", "err", err)
		return getInternalServerErrorResponse(codeInternalServerError, err.Error())
	}

	return &addResponse{
		Id: result,
	}
}

func (h *HTTP) DeletePostById(c *gin.Context) {
	resp := h.DeletePostByIdResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) DeletePostByIdResponse(c *gin.Context) response {
	id := c.Param("id")
	if id == "" {
		return getBadRequestWithMsgResponse(codeNoParam, "no param")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return getBadRequestWithMsgResponse(codeInvalidConvertion, err.Error())
	}

	err = h.service.DeletePostById(c.Request.Context(), idInt)
	if err != nil {
		if errors.Is(err, service.ErrNoSuchPost) {
			return getBadRequestWithMsgResponse(codeInvalidPostId, err.Error())
		}

		h.log.Errorw("Delete post by id.", "err", err)
		return getInternalServerErrorResponse(codeInternalServerError, err.Error())
	}

	return newOKResponse(codeOK)
}

func (h *HTTP) UpdatePost(c *gin.Context) {
	resp := h.UpdatePostResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) UpdatePostResponse(c *gin.Context) response {
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
