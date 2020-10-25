package routes

import (
	"github.com/gin-gonic/gin"
	"rain/internal/app/model"
	"rain/library/captcha"
	"rain/library/helper"
	"rain/library/response"
	"rain/library/token"
)

func LoginApi(c *gin.Context) {
	auth := model.Auth{}
	user, status := auth.Login(c)
	newToken, _ := token.CreateToken([]byte(helper.Env("SECRET_KEY")), c.GetHeader("Origin"), user.ID, true)

	if !status {
		resp.Error(c, 400, "error")
		return
	}
	token.ParseToken(newToken, []byte(helper.Env("SECRET_KEY")))
	resp.Success(c, "ok", gin.H{"user": user, "token": newToken})
}

func RegisterApi(c *gin.Context) {
	Id := c.Param("Id")
	VerifyValue := c.Param("VerifyValue")
	if captcha.VerifyCaptchaHandler(Id, VerifyValue) == false {
		resp.Error(c, 400, "验证码不正确")
		return
	}
	auth := model.Auth{}
	if (auth.Register(c, false)) == false {
		resp.Error(c, 400, "注册失败")
		return
	}
	resp.Success(c, "ok", "注册成功")
}
