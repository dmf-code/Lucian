package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": data})
	ctx.Abort()
}

func Fail(ctx *gin.Context, data interface{})  {
	ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "data": data})
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


