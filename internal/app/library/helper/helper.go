package helper

import (
	"app/library/config"
	"app/library/database"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Success(ctx *gin.Context, data interface{}, args ...interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": data, "args": args})
	ctx.Abort()
}

func Fail(ctx *gin.Context, data interface{})  {
	ctx.JSON(http.StatusOK, gin.H{"status": false, "data": data})
	ctx.Abort()
}

func Bad(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{"response": data})
}

// 丢弃BindJSON这种臃肿的获取值模式，采用灵活的MAP
func GetRequestJson(ctx *gin.Context) (requestMap map[string]interface{}) {
	requestData, err := ctx.GetRawData()
	if err != nil {
		Fail(ctx, "参数获取失败")
		return
	}
	err = json.Unmarshal(requestData, &requestMap)
	if err != nil {
		Fail(ctx, "参数获取失败")
	}

	fmt.Println(requestMap)
	return
}

func Db() (con *gorm.DB) {
	master := database.NewMySQL(config.RainDog).Master()
	return master.Write()
}

func Env(str string) (res string) {
	res = os.Getenv(str)
	return
}

// 判断某一个值是否含在切片之中
func InArray(need interface{}, haystack interface{}) bool {
	switch key := need.(type) {
	case int:
		for _, item := range haystack.([]int) {
			if item == key {
				return true
			}
		}
	case string:
		for _, item := range haystack.([]string) {
			if item == key {
				return true
			}
		}
	case int64:
		for _, item := range haystack.([]int64) {
			if item == key {
				return true
			}
		}
	case float64:
		for _, item := range haystack.([]float64) {
			if item == key {
				return true
			}
		}
	default:
		return false
	}
	return false
}

func Str2Uint(str string) (b uint) {
	a,_ := strconv.ParseUint(str, 10, 64)
	b = uint(a)
	return
}


func Float64ToInt(f float64) (res int) {
	tmp := strconv.FormatFloat(f, 'f', -1, 64)
	var err error
	res, err = strconv.Atoi(tmp)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	// 去除图片
	re, _ = regexp.Compile("\\<img[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func SubString(str, dot string, start, length int) string  {
	rs := []rune(str)
	end := start + length
	if end > len(rs) {
		end = len(rs)
	}
	subString := string(rs[start:end])
	if dot != "" {
		subString += dot
	}
	return subString
}
