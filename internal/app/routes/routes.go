package routes

import (
	"github.com/gin-gonic/gin"
	"rain/internal/app/middleware"
	"rain/library/captcha"
	"rain/library/go-fs"
	"rain/library/helper"
	"rain/library/uploader"
)

func Auth(r *gin.Engine) {
	r.POST("/login", LoginApi)
	r.POST("/register", RegisterApi)
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
		Front(front)
	}

	backend := r.Group("/backend")
	backend.Use(middleware.AccessTokenMiddleware())
	{

		backend.GET("ping", func(context *gin.Context) {
			helper.Success(context, "pong")
		})
		Backend(backend)
	}

	uploadInstance, err := uploader.New(r, uploader.GatherConfig{
		Path:      fs.StoragePath() + "upload",
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
