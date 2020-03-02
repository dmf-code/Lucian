package apis

import (
	"blog/model"
	"blog/utils/captcha"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)
type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginApi(c *gin.Context) {
	var loginInfo model.LoginInfo
	err := c.BindJSON(&loginInfo)
	if err != nil {
		fmt.Println(err)
		helper.Fail(c, 200,  "failed")
	}
	user, status := model.Login(loginInfo)
	fmt.Println(status)
	fmt.Println(user)
	if !status {
		helper.Fail(c, 200, "failed")
	}
	helper.Success(c, 200, gin.H{"user": user})
}

func RegisterApi(c *gin.Context) {
	var registerInfo model.RegisterInfo
	err := c.BindJSON(&registerInfo)

	if err != nil {
		fmt.Println(err)
	} else {
		if captcha.VerifyCaptchaHandler(registerInfo.Id, registerInfo.VerifyValue) == false {
			helper.Fail(c, 200,  "验证码不正确")
		}
		model.Register(registerInfo)
	}
	helper.Success(c, 200, gin.H{})
}
