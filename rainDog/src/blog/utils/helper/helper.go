package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
)

func Success(ctx *gin.Context,code int, data map[string]interface{}) {
	ctx.JSON(code, gin.H{"status": true, "data": data})
	ctx.Abort()
}

func Fail(ctx *gin.Context, code int, msg string)  {
	ctx.JSON(code, gin.H{"status":false, "msg": msg})
	ctx.Abort()
}



func Db(dbName string) (db *gorm.DB, err error) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	str := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, ip, port, dbName)
	db, err = gorm.Open("mysql", str)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Env(str string) (res string) {
	res = os.Getenv(str)
	return
}


