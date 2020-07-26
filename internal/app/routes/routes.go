package routes

import (
	"app/apis"
	"app/library/captcha"
	"app/library/helper"
	"app/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func Groups(r *gin.Engine) *gin.Engine {
	r.POST("/login", apis.LoginApi)
	r.POST("/register", apis.RegisterApi)
	r.POST("/getCaptcha", captcha.GenerateCaptchaHandler)
	//r.POST("/verifyCaptcha", captcha.VerifyCaptchaHandler)
	return r
}

func SetupRouter() *gin.Engine {

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r = Groups(r)
	front := r.Group("/front")
	{
		front.GET("ping", func(context *gin.Context) {
			helper.Success(context, "pong")
		})
		front.GET("/static/img", func(context *gin.Context) {
			workPath, _ := os.Getwd()
			dst := context.Query("url")
			imgPath := workPath + string(os.PathSeparator) + dst
			fmt.Println(imgPath)
			context.File(imgPath)
		})
		Front(front)
	}


	backend := r.Group("/backend")
	backend.Use(middleware.AccessTokenMiddleware())

	{

		backend.POST("/upload", func(context *gin.Context) {
			header, err := context.FormFile("file")
			if err != nil {
				//ignore
			}
			dst := header.Filename
			fmt.Println(os.Getwd())
			workPath, _ := os.Getwd()
			imgPath := "storages" + string(os.PathSeparator) + "upload" + string(os.PathSeparator) + dst
			path := workPath + string(os.PathSeparator) + imgPath
			// gin 简单做了封装,拷贝了文件流
			if err := context.SaveUploadedFile(header, path); err != nil {
				// ignore
			}

			helper.Success(context, gin.H{"path": imgPath})
		})

		backend.GET("ping", func(context *gin.Context) {
			helper.Success(context, "pong")
		})
		backend.GET("env", func(context *gin.Context) {
			envString := os.Environ()
			var envs map[string]string
			contains := []string{"APPDATA", "COMPUTERNAME", "GO111MODULE", "GOPATH", "GOPROXY", "GOROOT", "OS", "USER", "USERNAME"}
			envs = make(map[string]string)
			for i := 0; i < len(envString); i++ {
				tmp := strings.Split(envString[i], "=")
				if helper.InArray(tmp[0], contains) {
					envs[tmp[0]] = tmp[1]
				}
			}
			helper.Success(context, envs)
		})
		Backend(backend)
	}

	return r
}