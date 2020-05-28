package main

import (
	"app/bootstrap"
	"app/routes"
	"app/utils/mysqlTools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
)

func Migration() {
	bootstrap.InitTable()
}


func main() {

	// 配置日志
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 加载.env配置
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化Mysql连接池
	if !mysqlTools.GetInstance().InitDataPool() {
		log.Println("init database mysqlTools failure...")
		os.Exit(1)
	}

	// 迁移数据
	Migration()

	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println(secretKey)

	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}
