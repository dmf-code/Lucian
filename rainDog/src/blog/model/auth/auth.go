package auth

import (
	"blog/utils/helper"
	"blog/utils/model"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	model.BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) (user Users, status bool) {
	db := helper.Db("rain_dog")
	requestMap := helper.GetRequestJson(ctx)
	result := db.Table("user").
		Where("username = ? and password = ?", requestMap["username"], requestMap["password"]).
		First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return user, false
	}

	return user, true
}

func Register(ctx *gin.Context) (status bool) {
	db := helper.Db("rain_dog")
	requestMap := helper.GetRequestJson(ctx)
	if err := db.Table("user").Create(&requestMap).Error; err != nil {
		return false
	}

	return true
}
