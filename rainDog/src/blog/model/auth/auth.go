package auth

import (
	"blog/model/manage"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Login(ctx *gin.Context) (user manage.Admin, status bool) {
	db := helper.Db()
	requestMap := helper.GetRequestJson(ctx)
	result := db.Table("admin").
		Where("username = ? and password = ?", requestMap["username"], requestMap["password"]).
		First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return user, false
	}

	return user, true
}

func Register(ctx *gin.Context) (status bool) {
	db := helper.Db()
	requestMap := helper.GetRequestJson(ctx)
	if err := db.Table("admin").Create(&requestMap).Error; err != nil {
		return false
	}

	return true
}
