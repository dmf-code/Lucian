package auth

import (
	"blog/model/manage"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)


func Login(ctx *gin.Context) (user manage.Admin, status bool) {
	db := helper.Db()
	requestMap := helper.GetRequestJson(ctx)
	result := db.Table("admin").
		Where("username = ?", requestMap["username"]).
		First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return user, false
	}
	fmt.Println([]byte(user.Password))
	fmt.Println([]byte(requestMap["password"].(string)))
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestMap["password"].(string))); err != nil {
		fmt.Println(err)
		return user, false
	}

	return user, true
}

func Register(ctx *gin.Context) (status bool) {
	db := helper.Db()
	requestMap := helper.GetRequestJson(ctx)
	var err error
	requestMap["password"], err = bcrypt.GenerateFromPassword([]byte(requestMap["password"].(string)), bcrypt.DefaultCost)
	password := string(requestMap["password"].([]byte))
	username := requestMap["username"]
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(requestMap)

	if err = db.Table("admin").Create(&manage.Admin{Username: username.(string), Password: password}).Error; err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
