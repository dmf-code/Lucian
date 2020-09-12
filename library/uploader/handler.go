package uploader

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"rain/library/go-fs"
	"strconv"
	"strings"
)

const FIELD = "file"

var supportImageExtNames = []string{".jpg", ".jpeg", ".png", ".ico", ".svg", ".bmp", ".gif"}

/**
Handler the parse error
*/
func parseFormFail(context *gin.Context) {
	context.JSON(http.StatusBadRequest, gin.H{
		"message": "Can not parse form",
	})
}

/**
check a file is a image or not
*/
func isImage(extName string) bool {
	for i := 0; i < len(supportImageExtNames); i++ {
		if supportImageExtNames[i] == extName {
			return true
		}
	}
	return false
}

/**
Generate thumbnail
*/
func thumb(imagePath string) (outputPath string, err error) {
	var (
		file     *os.File
		img      image.Image
		filename = path.Base(imagePath)
	)

	extname := strings.ToLower(path.Ext(imagePath))

	outputPath = path.Join(Config.Path, Config.Image.Thumb.Path, filename)

	// 读取文件
	if file, err = os.Open(imagePath); err != nil {
		return
	}

	defer file.Close()

	// decode jpeg into image.Image
	switch extname {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
		break
	case ".png":
		img, err = png.Decode(file)
		break
	case ".gif":
		img, err = gif.Decode(file)
		break
	default:
		err = errors.New("Unsupport file type" + extname)
		return
	}

	if img == nil {
		err = errors.New("Generate thumbnail fail...")
		return
	}

	m := resize.Thumbnail(uint(Config.Image.Thumb.MaxWidth), uint(Config.Image.Thumb.MaxHeight), img, resize.Lanczos3)

	out, err := os.Create(outputPath)
	if err != nil {
		return
	}
	defer out.Close()

	// write new image to file

	//decode jpeg/png/gif into image.Image
	switch extname {
	case ".jpg", ".jpeg":
		jpeg.Encode(out, m, nil)
		break
	case ".png":
		png.Encode(out, m)
		break
	case ".gif":
		gif.Encode(out, m, nil)
		break
	default:
		err = errors.New("Unsupport file type" + extname)
		return
	}

	return
}

func UploadImage(context *gin.Context) {
	var (
		maxSize  = Config.Image.MaxSize // 最大上传大小
		distPath string                 // 最终的输出目录
		err      error
		file     *multipart.FileHeader
		src      multipart.File
		dist     *os.File
	)

	// source
	if file, err = context.FormFile(FIELD); err != nil {
		parseFormFail(context)
		return
	}

	extname := strings.ToLower(path.Ext(file.Filename))

	if isImage(extname) == false {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unsupport upload file type " + extname,
		})
		return
	}

	if file.Size > int64(maxSize) {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Upload file too large, The max upload limit is " + strconv.Itoa(int(maxSize)),
		})
		return
	}

	if src, err = file.Open(); err != nil {

	}

	defer src.Close()

	hash := md5.New()

	io.Copy(hash, src)

	md5string := hex.EncodeToString(hash.Sum([]byte("")))

	fileName := md5string + extname

	distPath = path.Join(Config.Path, Config.Image.Path, fileName)
	fmt.Println(distPath)
	if dist, err = os.Create(distPath); err != nil {

	}

	defer dist.Close()

	if src, err = file.Open(); err != nil {
		//
	}

	// Copy
	io.Copy(dist, src)

	// 压缩缩略图
	// 不管成功与否，都会进行下一步的返回
	if _, err := thumb(distPath); err != nil {

	}

	context.JSON(http.StatusOK, gin.H{
		"hash":     md5string,
		"filename": fileName,
		"origin":   file.Filename,
		"size":     file.Size,
	})
}

func UploadFile(context *gin.Context) {
	var (
		isSupportFile bool
		maxUploadSize = Config.Image.MaxSize  // 最大上传大小
		allowTypes    = Config.File.AllowType // 可上传的文件类型
		distPath      string                  // 最终的输出目录
		err           error
		file          *multipart.FileHeader
		src           multipart.File
		dist          *os.File
	)

	// Source
	if file, err = context.FormFile(FIELD); err != nil {
		parseFormFail(context)
		return
	}

	extname := path.Ext(file.Filename)

	if len(allowTypes) != 0 {
		for i := 0; i < len(allowTypes); i++ {
			if allowTypes[i] == extname {
				isSupportFile = true
				break
			}
		}

		if isSupportFile == false {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Unsupport upload file type " + extname,
			})
			return
		}
	}

	if file.Size > int64(maxUploadSize) {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Upload file too large, The max upload limit is " + strconv.Itoa(int(maxUploadSize)),
		})
		return
	}

	if src, err = file.Open(); err != nil {
		// open the file fail...
	}
	defer src.Close()

	hash := md5.New()

	io.Copy(hash, src)

	md5string := hex.EncodeToString(hash.Sum([]byte("")))

	fileName := md5string + extname

	// Destination
	distPath = path.Join(Config.Path, Config.File.Path, fileName)
	if dist, err = os.Create(distPath); err != nil {
		// create dist file fail...
	}
	defer dist.Close()

	// FIXME: open 2 times
	if src, err = file.Open(); err != nil {
		//
	}

	// Copy
	io.Copy(dist, src)

	context.JSON(http.StatusOK, gin.H{
		"hash":     md5string,
		"filename": fileName,
		"origin":   file.Filename,
		"size":     file.Size,
	})

}

/**
Get Origin image
*/
func GetOriginImage(context *gin.Context) {
	filename := context.Param("filename")
	originImagePath := path.Join(Config.Path, Config.Image.Path, filename)
	if fs.PathExists(originImagePath) == false {
		// if the path not found
		http.NotFound(context.Writer, context.Request)
		return
	}
	http.ServeFile(context.Writer, context.Request, originImagePath)
}

/**
Get thumb image
*/
func GetThumbImage(context *gin.Context) {
	filename := context.Param("filename")
	originImagePath := path.Join(Config.Path, Config.Image.Path, filename)
	thumbnailImagePath := path.Join(Config.Path, Config.Image.Thumb.Path, filename)
	if fs.PathExists(thumbnailImagePath) == false {
		// if thumbnail image not exist, try to get origin image
		if fs.PathExists(originImagePath) == true {
			http.ServeFile(context.Writer, context.Request, originImagePath)
			return
		}
		// if the path not found
		http.NotFound(context.Writer, context.Request)
		return
	}
	http.ServeFile(context.Writer, context.Request, thumbnailImagePath)
}

/**
Get file raw
*/
func GetFileRaw(context *gin.Context) {
	filename := context.Param("filename")
	filePath := path.Join(Config.Path, Config.File.Path, filename)
	if isExistFile := fs.PathExists(filePath); isExistFile == false {
		// if the path not found
		http.NotFound(context.Writer, context.Request)
		return
	}
	http.ServeFile(context.Writer, context.Request, filePath)
}

/**
Download a file
*/
func DownloadFile(context *gin.Context) {
	filename := context.Param("filename")
	filePath := path.Join(Config.Path, Config.File.Path, filename)
	if isExistFile := fs.PathExists(filePath); isExistFile == false {
		// if the path not found
		http.NotFound(context.Writer, context.Request)
		return
	}
	http.ServeFile(context.Writer, context.Request, filePath)
}
