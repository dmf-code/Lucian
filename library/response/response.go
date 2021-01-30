package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Msg	 string `json:"msg"`
	Data interface{} `json:"data"`
}

func Error(ctx *gin.Context, code int, msg string, data ...interface{})  {
	response(ctx, code, msg, data...)
}

func Success(ctx *gin.Context, msg string, data ...interface{})  {
	response(ctx, 0, msg, data...)
}

func Paginate(ctx *gin.Context, msg string, items interface{}, total int, page int, pageSize int) {

	if page == 0 {
		page = 1
	}

	switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
	}

	paginatorMap := make(map[string]interface{})
	paginatorMap["items"] = items
	paginatorMap["total"] = total
	paginatorMap["page"] = page
	paginatorMap["page_size"] = pageSize
	response(ctx, 0, msg, paginatorMap)
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
	ctx.Abort()
}
