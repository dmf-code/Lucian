package uploader

import (
	"app/library/go-fs"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"path"
)

type FileConfig struct {
	Path		string		`valid:"required,length(1|20)"` 	// 普通文件的存放目录
	MaxSize 	int			`valid:"required"`					// 普通文件上传的限制大小，单位byte, 最大单位1GB
	AllowType	[]string	`valid:"required,length(0|100)"`	// 允许上传的文件后缀名
}

type ThumbConfig struct {
	Path      string `valid:"required,length(1|20)"` // 缩略图存放路径
	MaxWidth  int    `valid:"required"`              // 缩略图最大宽度
	MaxHeight int    `valid:"required"`              // 缩略图最大高度
}

type ImageConfig struct {
	Path    string `valid:"required,length(1|20)"` // 图片存储路径
	MaxSize int    `valid:"required"`              // 最大图片上传限制，单位byte
	Thumb 	ThumbConfig
}

type GatherConfig struct {
	Path      string `valid:"required,length(1|100)"` //文件上传的根目录
	UrlPrefix string `valid:"required,length(0|20)"` // api的url前缀
	File      FileConfig
	Image     ImageConfig
}

type Uploader struct {
	Upload   *gin.RouterGroup
	Download *gin.RouterGroup
	Config   GatherConfig
	Engine   *gin.Engine
}

var Config	GatherConfig

func InitUploader(c *GatherConfig) (err error) {

	if err = fs.IsDir(c.Path); err != nil {
		return
	}

	if err = fs.IsDir(path.Join(c.Path, c.File.Path)); err != nil {
		return
	}

	if err = fs.IsDir(path.Join(c.Path, c.Image.Path)); err != nil {
		return
	}

	if err = fs.IsDir(path.Join(c.Path, c.Image.Thumb.Path)); err != nil {
		return
	}
	return
}

func New(e *gin.Engine, c GatherConfig) (u *Uploader, err error)  {
	Config = c
	var (
		isValidConfig bool
	)

	if isValidConfig, err = govalidator.ValidateStruct(c); err != nil {
		err = errors.New(`invalid uploader config: [` + err.Error() + `]`)
		return
	} else {
		if isValidConfig == false {
			err = errors.New("invalid Uploader config")
			return
		}
	}

	if err = InitUploader(&Config); err != nil {
		return
	}

	uploader := e.Group(Config.UrlPrefix + "/upload")

	downloader := e.Group(Config.UrlPrefix + "/download")

	downloader.Use(func(context *gin.Context) {
		header := context.Writer.Header()

		// alone dns prefect
		header.Set("X-DNS-Prefetch-Control", "on")
		// IE No Open
		header.Set("X-Download-Options", "noopen")
		// not cache
		header.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		header.Set("Expires", "max-age=0")
		// Content Security Policy
		header.Set("Content-Security-Policy", "default-src 'self'")
		// xss protect
		// it will caught some problems is old IE
		header.Set("X-XSS-Protection", "1; mode=block")
		// Referrer Policy
		header.Set("Referrer-Header", "no-referrer")
		// cros frame, allow same origin
		header.Set("X-Frame-Options", "SAMEORIGIN")
		// HSTS
		header.Set("Strict-Transport-Security", "max-age=5184000;includeSubDomains")
		// no sniff
		header.Set("X-Content-Type-Options", "nosniff")
	})

	return &Uploader{
		Upload: uploader,
		Download: downloader,
		Config: c,
		Engine: e,
	}, nil
}



func (u *Uploader) Resolve() {
	// upload the file/image
	u.Upload.POST("/image", UploadImage)
	u.Upload.POST("/file", UploadFile)

	// get file which upload
	uploadFile := u.Download.Group("/file")
	uploadFile.GET("/raw/:filename", GetFileRaw)
	uploadFile.GET("/download/:filename", DownloadFile)

	// get image which upload
	downloadImage := u.Download.Group("/image")
	downloadImage.GET("/thumbnail/:filename", GetThumbImage)
	downloadImage.GET("/origin/:filename", GetOriginImage)

	return
}
