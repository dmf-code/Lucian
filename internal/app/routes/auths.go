package routes

import (
	"github.com/gin-gonic/gin"
	"rain/internal/app/model"
	"rain/library/captcha"
	"rain/library/helper"
	"rain/library/token"
)

func LoginApi(c *gin.Context) {
	auth := model.Auth{}
	user, status := auth.Login(c)
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
		helper.Fail(c, "验证码不正确")
		return
	}
	auth := model.Auth{}
	if (auth.Register(c, false)) == false {
		helper.Fail(c, "注册失败")
		return
	}
	helper.Success(c, "注册成功")
}
