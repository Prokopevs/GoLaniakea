package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type addResponse struct {
	Id int `json:"id"`
}

func (a *addResponse) writeJSON(c *gin.Context) {
	c.JSON(http.StatusOK, a)
}

type OKStruct struct {
	jsonData interface{}
}

func newOKResponse(data interface{}) *OKStruct {
	return &OKStruct{
		jsonData: data,
	}
}

func (p *OKStruct) writeJSON(c *gin.Context) {
	c.JSON(http.StatusOK, p.jsonData)
}
