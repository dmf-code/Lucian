package routes

import (
	"blog/apis"
	"blog/utils/captcha"
	"github.com/gin-gonic/gin"
)

func Groups(r *gin.Engine) *gin.Engine {
	r.POST("/login", apis.LoginApi)
	r.POST("/register", apis.RegisterApi)
	r.POST("/getCaptcha", captcha.GenerateCaptchaHandler)
	//r.POST("/verifyCaptcha", captcha.VerifyCaptchaHandler)
	return r
}
