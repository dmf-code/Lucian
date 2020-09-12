package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"rain/internal/app/bootstrap"
	"rain/internal/app/routes"
	"rain/library/permission"
)

func main() {
	// 配置日志
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 加载.env配置
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 权限初始化
	permission.Init()

	// 迁移数据
	bootstrap.InitTable()

	r, err := routes.SetupRouter()

	if err != nil {
		panic("初始化路由失败： " + err.Error())
	}

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}
