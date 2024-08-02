package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status    int    `json:"-"`
	Code      Code   `json:"code"`
	ErrorInfo string `json:"errorInfo"`
}

func (e errorResponse) writeJSON(c *gin.Context) {
	c.JSON(e.Status, e)
}

func getBadRequestWithMsgResponse(code Code, msg string) errorResponse {
	return errorResponse{
		Status:    http.StatusBadRequest,
		Code:      code,
		ErrorInfo: msg,
	}
}

func getInternalServerErrorResponse(code Code, msg string) errorResponse {
	return errorResponse{
		Status:    http.StatusInternalServerError,
		Code:      code,
		ErrorInfo: msg,
	}
}

func getForbiddenRequestWithMsgResponse(code Code, msg string) errorResponse {
	return errorResponse{
		Status:    http.StatusForbidden,
		Code:      code,
		ErrorInfo: msg,
	}
}
