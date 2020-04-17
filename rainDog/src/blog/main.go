package main

import (
	"blog/middleware"
	"blog/model/manage"
	"blog/routes"
	"blog/utils/helper"
	"blog/utils/mysqlTools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
)

func setupRouter() *gin.Engine {

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r = routes.Groups(r)

	front := r.Group("/front")
	{
		front.GET("ping", func(context *gin.Context) {
			helper.Success(context, "pong")
		})
		routes.Front(front)
	}


	backend := r.Group("/backend")
	backend.Use(middleware.AccessTokenMiddleware())

	{
		backend.GET("ping", func(context *gin.Context) {
			helper.Success(context, "pong")
		})
		routes.Backend(backend)
	}

	return r
}

func Migration() {
	table()
	//row()
}

func table() {
	fmt.Println(mysqlTools.GetInstance().GetMysqlDB().AutoMigrate(new(manage.Menu)).Error)
	fmt.Println(mysqlTools.GetInstance().GetMysqlDB().AutoMigrate(new(manage.Role)).Error)
	fmt.Println(mysqlTools.GetInstance().GetMysqlDB().AutoMigrate(new(manage.Admin)).Error)
	fmt.Println(mysqlTools.GetInstance().GetMysqlDB().AutoMigrate(new(manage.RoleMenu)).Error)
	fmt.Println(mysqlTools.GetInstance().GetMysqlDB().AutoMigrate(new(manage.AdminRole)).Error)
	if os.Getenv("INIT_ADMIN_TABLE") == "true" {
		row()
	}
}

func row() {
	var father, child manage.Menu
	helper.Db().Create(manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/refresh", Name: "refresh", Sequence: 5, Type: 4, Component: "@/views/Refresh", Meta: "", Icon: "", OperateType: "none"})

	father = manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/", Name: "index", Sequence: 5, Type: 4, Component: "@/views/front/Index", Meta: "", Icon: "el-icon-s-home", OperateType: "view"}
	helper.Db().Create(&father)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "/", Name: "home", Sequence: 5, Type: 4, Component: "@/views/front/Home", Meta: "", Icon: "el-icon-s-home", OperateType: "view"}
	helper.Db().Create(&child)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "article/:id", Name: "article", Sequence: 5, Type: 4, Component: "@/views/front/pages/article/Index", Meta: "", Icon: "el-icon-s-home", OperateType: "view"}
	helper.Db().Create(&child)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "docs/:path", Name: "docs", Sequence: 5, Type: 4, Component: "@/views/front/pages/doc/Index", Meta: "", Icon: "el-icon-s-home", OperateType: "view"}
	helper.Db().Create(&child)

	helper.Db().Create(manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/login", Name: "login", Sequence: 5, Type: 4, Component: "@/views/auth/Login", Meta: "", Icon: "", OperateType: "view"})

	helper.Db().Create(manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/register", Name: "register", Sequence: 5, Type: 4, Component: "@/views/auth/Register", Meta: "", Icon: "", OperateType: "view"})

	father = manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/admin", Name: "admin", Sequence: 5, Type: 1, Component: "@/views/admin/Index", Meta: "requireAuth:true;", Icon: "el-icon-s-home", OperateType: "view"}
	helper.Db().Create(&father)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "dashboard", Name: "首页", Sequence: 5, Type: 2, Component: "@/views/admin/Dashboard", Meta: "requireAuth:true;", Icon: "el-icon-s-home", OperateType: "view"}
	helper.Db().Create(&child)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "menu", Name: "菜单", Sequence: 5, Type: 2, Component: "@/views/admin/Dashboard", Meta: "requireAuth:true;", Icon: "el-icon-s-home", OperateType: "view"}
	helper.Db().Create(&child)

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

	// 初始化数据表
	table()

	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println(secretKey)

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}
