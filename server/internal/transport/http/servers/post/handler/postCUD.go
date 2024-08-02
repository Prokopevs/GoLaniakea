package handler

import (
	"errors"
	"strconv"

	"github.com/Prokopevs/GoLaniakea/server/internal/service"
	"github.com/gin-gonic/gin"
)

type PostJSON struct {
	Id           int    `json:"id"`
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
	password := c.Request.Header.Get("password")
	if password == "" || password != h.password {
		return getForbiddenRequestWithMsgResponse(codeHeaderErr, "header error")
	}

	p := &PostJSON{}
	if err := c.ShouldBindJSON(p); err != nil {
		return getBadRequestWithMsgResponse(codeBadBody, err.Error())
	}

	code, err := p.validate("create")
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
	password := c.Request.Header.Get("password")
	if password == "" || password != h.password {
		return getForbiddenRequestWithMsgResponse(codeHeaderErr, "header error")
	}

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
	password := c.Request.Header.Get("password")
	if password == "" || password != h.password {
		return getForbiddenRequestWithMsgResponse(codeHeaderErr, "header error")
	}

	p := &PostJSON{}
	if err := c.ShouldBindJSON(p); err != nil {
		return getBadRequestWithMsgResponse(codeBadBody, err.Error())
	}

	code, err := p.validate("update")
	if err != nil {
		return getBadRequestWithMsgResponse(code, err.Error())
	}

	post, err := convertPostJSONToPost(p)
	if err != nil {
		return getBadRequestWithMsgResponse(codeInvalidConvertion, err.Error())
	}

	err = h.service.UpdatePost(c.Request.Context(), post)
	if err != nil {
		if errors.Is(err, service.ErrNoSuchPost) {
			return getBadRequestWithMsgResponse(codeInvalidPostId, err.Error())
		}

		h.log.Errorw("Delete post by id.", "err", err)
		return getInternalServerErrorResponse(codeInternalServerError, err.Error())
	}

	return newOKResponse(codeOK)
}
