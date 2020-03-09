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
	Id string `json:"id"`
	VerifyValue string `json:"verify_value"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id string `json:"id"`
	VerifyValue string `json:"verify_value"`
}

type Users struct {
	model.BaseModel
	Username string
	password string
}

func Login(info LoginInfo) (user Users, status bool) {
	db, err := helper.Db("rain_dog")

	if err != nil {
		fmt.Println(err)
		return user, false
	}

	result := db.Table("user").
		Where("username = ? and password = ?", info.Username, info.Password).
		First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return user, false
	}

	defer db.Close()
	return user, true
}

func Register(info RegisterInfo) (status bool) {
	db, err := helper.Db("rain_dog")

	if err != nil {
		fmt.Println(err)
		return false
	}

	id := db.Table("user").Create(&info)

	fmt.Println(id)

	defer db.Close()
	return true
}
