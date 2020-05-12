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
	"strings"
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
		backend.GET("env", func(context *gin.Context) {
			envString := os.Environ()
			var envs map[string]string
			contains := []string{"APPDATA", "COMPUTERNAME", "GO111MODULE", "GOPATH", "GOPROXY", "GOROOT", "OS", "USER", "USERNAME"}
			envs = make(map[string]string)
			for i := 0; i < len(envString); i++ {
				tmp := strings.Split(envString[i], "=")
				if helper.InArray(tmp[0], contains) {
					envs[tmp[0]] = tmp[1]
				}
			}
			helper.Success(context, envs)
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
	db := mysqlTools.GetInstance().GetMysqlDB().Set("gorm:table_options", "ENGINE=InnoDB")
	fmt.Println(db.AutoMigrate(new(manage.Menu)).Error)
	fmt.Println(db.AutoMigrate(new(manage.Role)).Error)
	fmt.Println(db.AutoMigrate(new(manage.Admin)).Error)
	fmt.Println(db.AutoMigrate(new(manage.RoleMenu)).Error)
	fmt.Println(db.AutoMigrate(new(manage.AdminRole)).Error)
	if os.Getenv("INIT_ADMIN_TABLE") == "true" {
		row()
	}
}

func row() {
	var grandfather, father, child, unit manage.Menu

	unit = manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/refresh", Name: "refresh", Sequence: 5, Type: 4, Component: "@/views/Refresh",  Icon: "", OperateType: "none"}
	helper.Db().Create(&unit)

	unit = manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/login", Name: "login", Sequence: 5, Type: 4, Component: "@/views/auth/Login",  Icon: "", OperateType: "view"}
	helper.Db().Create(&unit)

	unit = manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/register", Name: "register", Sequence: 5, Type: 4, Component: "@/views/auth/Register",  Icon: "", OperateType: "view"}
	helper.Db().Create(&unit)

	// 前端
	father = manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/", Name: "index", Sequence: 5, Type: 4, Component: "@/views/front/Index",  Icon: "", OperateType: "view"}
	helper.Db().Create(&father)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "/", Name: "home", Sequence: 5, Type: 4, Component: "@/views/front/Home",  Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "article/:id", Name: "article", Sequence: 5, Type: 4, Component: "@/views/front/pages/article/Index",  Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "docs/:path", Name: "docs", Sequence: 5, Type: 4, Component: "@/views/front/pages/doc/Index",  Icon: "", OperateType: "view"}
	helper.Db().Create(&child)


	grandfather = manage.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/admin", Name: "admin", Sequence: 5, Type: 1, Component: "@/views/admin/Index", Icon: "", OperateType: "view"}
	helper.Db().Create(&grandfather)

	father = manage.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "dashboard", Name: "首页", Sequence: 5, Type: 2, Component: "@/views/admin/Dashboard", Icon: "", OperateType: "view"}
	helper.Db().Create(&father)

	// 权限管理
	father = manage.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "", Name: "权限管理", Sequence: 5, Type: 1, Component: "",  Icon: "el-icon-s-operation", OperateType: "none"}
	helper.Db().Create(&father)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "menu", Name: "菜单", Sequence: 5, Type: 2, Component: "@/views/admin/pages/menu/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "role", Name: "角色", Sequence: 5, Type: 2, Component: "@/views/admin/pages/role/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	// 文章管理
	father = manage.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "", Name: "文章管理", Sequence: 5, Type: 1, Component: "",  Icon: "el-icon-s-data", OperateType: "none"}
	helper.Db().Create(&father)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "article", Name: "列表", Sequence: 5, Type: 2, Component: "@/views/admin/pages/article/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "addArticle", Name: "添加", Sequence: 5, Type: 4, Component: "@/views/admin/pages/article/Add", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	// 分类管理
	father = manage.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "", Name: "分类管理", Sequence: 5, Type: 1, Component: "",  Icon: "el-icon-menu", OperateType: "none"}
	helper.Db().Create(&father)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "category", Name: "列表", Sequence: 5, Type: 2, Component: "@/views/admin/pages/category/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "addCategory", Name: "添加", Sequence: 5, Type: 4, Component: "@/views/admin/pages/category/Add", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	// 标签管理
	father = manage.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "", Name: "标签管理", Sequence: 5, Type: 1, Component: "",  Icon: "el-icon-collection-tag", OperateType: "none"}
	helper.Db().Create(&father)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "tag", Name: "列表", Sequence: 5, Type: 2, Component: "@/views/admin/pages/tag/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = manage.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "addTag", Name: "添加", Sequence: 5, Type: 4, Component: "@/views/admin/pages/tag/Add", Icon: "", OperateType: "view"}
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
