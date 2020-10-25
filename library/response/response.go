package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Msg	 string
	Data interface{}
}

func Error(ctx *gin.Context, code int, msg string, data ...interface{})  {
	response(ctx, code, msg, data...)
}

func Success(ctx *gin.Context, code int, msg string, data ...interface{})  {
	response(ctx, 0, msg, data...)
}

func response(ctx *gin.Context, code int, msg string, data ...interface{}) {
	resp := Response{
		Code: code,
		Msg: msg,
		Data: data,
	}

	if len(data) == 1 {
		resp.Data = data[0]
	}

	ctx.JSON(http.StatusOK, resp)
}
