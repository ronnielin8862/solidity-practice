package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

const (
	Success = 0
	Fail    = 100
)

func RespOk(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: Success,
		Msg:  msg,
		Data: data,
	})
}

func RespFail(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: Fail,
		Msg:  msg,
		Data: data,
	})
}
