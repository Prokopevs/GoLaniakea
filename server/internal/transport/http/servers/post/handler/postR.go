package handler

import (
	"errors"
	"strconv"

	"github.com/Prokopevs/GoLaniakea/server/internal/service"
	"github.com/gin-gonic/gin"
)

func (h *HTTP) GetPosts(c *gin.Context) {
	resp := h.GetPostsResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) GetPostsResponse(c *gin.Context) response {
	category := c.Query("category")
	page := c.Query("page")
	limit := c.Query("limit")

	categoryInt, pageInt, limitInt, err := updateGetParams(category, page, limit)
	if err != nil {
		return getBadRequestWithMsgResponse(codeInvalidConvertion, err.Error())
	}

	res, err := h.service.GetPosts(c.Request.Context(), categoryInt, pageInt, limitInt)
	if err != nil {
		h.log.Errorw("Get posts.", "err", err)
		return getInternalServerErrorResponse(codeInternalServerError, err.Error())
	}

	return newOKResponse(convertRankPostsToPosts(res))
}

func (h *HTTP) GetPostById(c *gin.Context) {
	resp := h.GetPostByIdResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) GetPostByIdResponse(c *gin.Context) response {
	id := c.Param("id")
	if id == "" {
		return getBadRequestWithMsgResponse(codeNoParam, "no param")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return getBadRequestWithMsgResponse(codeInvalidConvertion, "invalid param")
	}

	res, err := h.service.GetPostById(c.Request.Context(), idInt)
	if err != nil {
		if errors.Is(err, service.ErrNoSuchPost) {
			return getBadRequestWithMsgResponse(codeInvalidPostId, err.Error())
		}

		h.log.Errorw("Get post by id.", "err", err)
		return getInternalServerErrorResponse(codeInternalServerError, err.Error())
	}

	return newOKResponse(res)
}

func (h *HTTP) GetTotalCount(c *gin.Context) {
	resp := h.GetTotalCountResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) GetTotalCountResponse(c *gin.Context) response {
	res, err := h.service.GetTotalCount(c.Request.Context())
	if err != nil {
		h.log.Errorw("Get total count.", "err", err)
		return getInternalServerErrorResponse(codeInternalServerError, err.Error())
	}

	return newOKResponse(res)
}

func (h *HTTP) GetInteresting(c *gin.Context) {
	resp := h.GetInterestingResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) GetInterestingResponse(c *gin.Context) response {
	res, err := h.service.GetInteresting(c.Request.Context())
	if err != nil {
		h.log.Errorw("Get interesting.", "err", err)
		return getInternalServerErrorResponse(codeInternalServerError, err.Error())
	}

	return newOKResponse(convertPostsToInteresting(res))
}
