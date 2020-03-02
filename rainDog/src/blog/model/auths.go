package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	gorm.Model
	Username string
	password string
}

func Login(info LoginInfo) (Users, bool) {
	db, err := gorm.Open("mysql", "root:root@/rain_dog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	var user Users
	result := db.Where("username = ? and password = ?", info.Username, info.Password).First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return user, false
	}

	fmt.Println(user)

	defer db.Close()
	return user, true
}

func Register(info RegisterInfo) bool {
	db, err := gorm.Open("mysql", "root:root@(192.168.3.9:9003)/rain_dog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	id := db.Table("user").Create(&info)

	fmt.Println(id)

	defer db.Close()
	return true
}
