package cmd

import (

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"rain/internal/app/model"
	"rain/library/helper"
)

var MigrationCmd = &cobra.Command{
	Use:   "Migration",
	Short: "系统数据初始化",
	Run: func(cmd *cobra.Command, args []string) {
		initTableData()
	},
}

func initTableData() {
	var grandfather, father, child, unit model.Menu

	unit = model.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/refresh", Name: "refresh", Sequence: 5, Type: 4, Component: "views/Refresh", Icon: "", OperateType: "none"}
	helper.Db().Create(&unit)

	unit = model.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/login", Name: "login", Sequence: 5, Type: 4, Component: "views/auth/Login", Icon: "", OperateType: "view"}
	helper.Db().Create(&unit)

	unit = model.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/register", Name: "register", Sequence: 5, Type: 4, Component: "views/auth/Register", Icon: "", OperateType: "view"}
	helper.Db().Create(&unit)

	// 前端
	father = model.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/", Name: "index", Sequence: 5, Type: 4, Component: "views/front/Index", Icon: "", OperateType: "view"}
	helper.Db().Create(&father)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "/", Name: "home", Sequence: 5, Type: 4, Component: "views/front/Home", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "article/:id", Name: "article", Sequence: 5, Type: 4, Component: "views/front/pages/article/Index", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "docs/:path", Name: "docs", Sequence: 5, Type: 4, Component: "views/front/pages/doc/Index", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	grandfather = model.Menu{Status: 1, Memo: "", ParentID: 0, Url: "/admin", Name: "admin", Sequence: 5, Type: 1, Component: "views/admin/Index", Icon: "", OperateType: "view"}
	helper.Db().Create(&grandfather)

	father = model.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "dashboard", Name: "首页", Sequence: 5, Type: 2, Component: "views/admin/Dashboard", Icon: "", OperateType: "view"}
	helper.Db().Create(&father)

	// 权限管理
	father = model.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "", Name: "权限管理", Sequence: 5, Type: 1, Component: "views/admin/Route", Icon: "el-icon-s-operation", OperateType: "none"}
	helper.Db().Create(&father)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "user", Name: "用户", Sequence: 5, Type: 2, Component: "views/admin/pages/user/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "menu", Name: "菜单", Sequence: 5, Type: 2, Component: "views/admin/pages/menu/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "role", Name: "角色", Sequence: 5, Type: 2, Component: "views/admin/pages/role/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	// 文章管理
	father = model.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "", Name: "文章管理", Sequence: 5, Type: 1, Component: "views/admin/Route", Icon: "el-icon-s-data", OperateType: "none"}
	helper.Db().Create(&father)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "article", Name: "列表", Sequence: 5, Type: 2, Component: "views/admin/pages/article/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "addArticle", Name: "添加", Sequence: 5, Type: 4, Component: "views/admin/pages/article/Add", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	// 分类管理
	father = model.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "", Name: "分类管理", Sequence: 5, Type: 1, Component: "views/admin/Route", Icon: "el-icon-menu", OperateType: "none"}
	helper.Db().Create(&father)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "category", Name: "列表", Sequence: 5, Type: 2, Component: "views/admin/pages/category/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "addCategory", Name: "添加", Sequence: 5, Type: 4, Component: "views/admin/pages/category/Add", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	// 标签管理
	father = model.Menu{Status: 1, Memo: "", ParentID: uint64(grandfather.ID), Url: "", Name: "标签管理", Sequence: 5, Type: 1, Component: "views/admin/Route", Icon: "el-icon-collection-tag", OperateType: "none"}
	helper.Db().Create(&father)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "tag", Name: "列表", Sequence: 5, Type: 2, Component: "views/admin/pages/tag/List", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	child = model.Menu{Status: 1, Memo: "", ParentID: uint64(father.ID), Url: "addTag", Name: "添加", Sequence: 5, Type: 4, Component: "views/admin/pages/tag/Add", Icon: "", OperateType: "view"}
	helper.Db().Create(&child)

	// 管理员账号生成
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	account := model.Admin{Username: "admin", Password: string(password)}
	helper.Db().Create(&account)

	// 管理员角色生成
	role := model.Role{Name: "super_admin", Sequence: 5, Memo: "超级管理员"}
	helper.Db().Create(&role)

	// 账号角色关联
	account2role := model.AdminRole{AdminId: uint64(account.ID), RoleId: uint64(role.ID)}
	helper.Db().Create(&account2role)
}

