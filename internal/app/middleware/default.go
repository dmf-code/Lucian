package middleware

import (
	"app/bootstrap/Table"
	"app/utils/helper"
	"app/utils/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccessTokenMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authored :=c.Request.Header.Get("token")
		fmt.Println(authored)
		if data, err := token.ParseToken(authored, []byte(helper.Env("SECRET_KEY"))); err == nil {
			// 验证通过，会继续访问下一个中间件
			var user Table.Admin
			db := helper.Db()
			uid := token.GetIdFromClaims("uid", data)
			fmt.Println(data)
			db.Table("admin").Where("id = ?", uid).First(&user)
			fmt.Println(user)
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"访问未授权"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
	}
}
