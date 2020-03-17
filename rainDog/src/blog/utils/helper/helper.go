package helper

import (
	"blog/utils/mysqlTools"
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



func Db(dbName string) (con *gorm.DB) {
	return mysqlTools.GetInstance().GetMysqlDB()
}

func Env(str string) (res string) {
	res = os.Getenv(str)
	return
}


