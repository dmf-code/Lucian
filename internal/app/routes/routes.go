package routes

import (
	"app/apis"
	"app/library/captcha"
	"app/library/go-fs"
	"app/library/helper"
	"app/library/uploader"
	"app/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func Auth(r *gin.Engine) {
	r.POST("/login", apis.LoginApi)
	r.POST("/register", apis.RegisterApi)
	r.POST("/getCaptcha", captcha.GenerateCaptchaHandler)
	//r.POST("/verifyCaptcha", captcha.VerifyCaptchaHandler)
}

func SetupRouter() (e *gin.Engine, err error) {

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	Auth(r)
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

	uploadInstance, err := uploader.New(r, uploader.GatherConfig{
		Path: fs.StoragePath() + "upload",
		UrlPrefix: "/common",
		File: uploader.FileConfig{
			Path:      "files",
			MaxSize:   10485760,
			AllowType: []string{".xls", ".txt"},
		},
		Image: uploader.ImageConfig{
			Path:    "images",
			MaxSize: 10485760,
			Thumb: uploader.ThumbConfig{
				Path:      "thumb",
				MaxWidth:  300,
				MaxHeight: 300,
			},
		},
	})

	if err != nil {
		return &gin.Engine{}, err
	}

	uploadInstance.Resolve()

	return r, nil
}