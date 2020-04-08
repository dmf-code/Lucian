package middleware

import (
	"blog/utils/helper"
	"blog/utils/permission"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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

// 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// 检查请求路径是否包含指定前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(ctx *gin.Context) bool {
		path := ctx.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}

		return true
	}
}

// 包含则不跳过
func AllowPathPrefixNotSkipper(prefixes ...string) SkipperFunc {
	return func(ctx *gin.Context) bool {
		path := JoinRouter(ctx.Request.Method, ctx.Request.URL.Path)
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// 拼接路由
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}

	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}
