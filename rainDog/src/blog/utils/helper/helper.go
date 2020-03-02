package helper

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context,code int, data map[string]interface{}) {

	ctx.JSON(code, gin.H{"status": true, "data": data})
}

func Fail(ctx *gin.Context, code int, msg string)  {
	ctx.JSON(code, gin.H{"status":false, "msg": msg})
}

