package middleware

import (
	"blog/utils/helper"
	"blog/utils/permission"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 权限中间间
func CasbinMiddleware(userId int, path string, method string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if b, err := permission.CheckPermission(strconv.Itoa(userId), path, method); err != nil {
			fmt.Println(err)
			helper.Fail(ctx, "error")
			return
		} else if !b {
			helper.Fail(ctx, "没有访问权限")
			return
		}
		ctx.Next()
	}
}
