package apis

import (
	"blog/model/auth"
	"blog/utils/captcha"
	"blog/utils/helper"
	"blog/utils/token"
	"fmt"
	"github.com/gin-gonic/gin"
)
type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginApi(c *gin.Context) {
	user, status := auth.Login(c)
	fmt.Println(status)
	fmt.Println(user)
	newToken, _ := token.CreateToken([]byte(helper.Env("SECRET_KEY")), c.GetHeader("Origin"), user.ID, true)

	if !status {
		helper.Fail(c, "failed")
		return
	}
	token.ParseToken(newToken, []byte(helper.Env("SECRET_KEY")))
	helper.Success(c, gin.H{"user": user, "token": newToken})
}

func RegisterApi(c *gin.Context) {
	Id := c.Param("Id")
	VerifyValue := c.Param("VerifyValue")
	if captcha.VerifyCaptchaHandler(Id, VerifyValue) == false {
			helper.Fail(c,  "验证码不正确")
	}
	auth.Register(c)
	helper.Success(c, "注册成功")
}
