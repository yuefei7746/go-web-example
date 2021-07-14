package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	MsgOK = "ok"
)

var (
	OK = ResponseEntity{
		Code: http.StatusOK,
		Msg:  MsgOK,
	}
)

type ResponseEntity struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseOK(data interface{}) *ResponseEntity {
	return &ResponseEntity{
		Code: http.StatusOK,
		Msg:  MsgOK,
		Data: data,
	}
}

func ResponseError(errMsg string) *ResponseEntity {
	return &ResponseEntity{
		Code: http.StatusInternalServerError,
		Msg:  errMsg,
	}
}

func renderOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseOK(data))
}

func renderError(c *gin.Context, errMsg string) {
	c.JSON(http.StatusInternalServerError, ResponseError(errMsg))
}
