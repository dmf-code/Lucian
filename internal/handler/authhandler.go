package handler

import (
	"github.com/gin-gonic/gin"
	"rain/internal/model"
	"rain/library/captcha"
	"rain/library/helper"
	resp "rain/library/response"
	"rain/library/token"
)

func Auth(r *gin.Engine) {
	r.POST("/login", loginAuth)
	r.POST("/register", registerAuth)
	r.POST("/getCaptcha", captcha.GenerateCaptchaHandler)
	//r.POST("/verifyCaptcha", captcha.VerifyCaptchaHandler)
}

func loginAuth(c *gin.Context) {
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

func registerAuth(c *gin.Context) {
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

