package auth

import (
	"blog/utils/helper"
	"blog/utils/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type RegisterInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users struct {
	model.BaseModel
	Username string
	password string
}

func Login(info LoginInfo) (user Users, status bool) {
	db := helper.Db("rain_dog")

	result := db.Table("user").
		Where("username = ? and password = ?", info.Username, info.Password).
		First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return user, false
	}

	return user, true
}

func Register(info RegisterInfo) (status bool) {
	db := helper.Db("rain_dog")

	if err := db.Table("user").Create(&info).Error; err != nil {
		return false
	}

	return true
}
