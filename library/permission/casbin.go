package permission

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"rain/internal/model"
	"rain/library/helper"
	"strconv"
	"strings"
)

const (
	PrefixUserId = "u"
	PrefixRoleId = "r"
)

var Enforcer *casbin.Enforcer

func Init() {

	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	a, err := gormadapter.NewAdapter("mysql", os.Getenv("DSN"))

	if err != nil {
		panic(err)
		return
	}
	fmt.Println(workPath)
	enforcer, err := casbin.NewEnforcer(workPath+string(os.PathSeparator)+"config/rbac_model.conf", a)

	if err != nil {
		panic(err)
		return
	}
	enforcer.EnableLog(true)
	var roles []model.Role
	db := helper.Db()
	if err = db.Table("role").Find(&roles).Error; err != nil {
		fmt.Println(err)
	}

	for _, role := range roles {
		setRolePermission(db, enforcer, uint64(role.ID))
	}

	if err = enforcer.LoadPolicy(); err != nil {
		return
	}

	Enforcer = enforcer
}

// 设置角色权限
func setRolePermission(db *gorm.DB, enforcer *casbin.Enforcer, roleId uint64) {
	var roleMenus []model.RoleMenu
	if err := db.Model(&model.RoleMenu{RoleId: roleId}).Find(&roleMenus).Error; err != nil {
		fmt.Println(err)
	}
	for _, roleMenu := range roleMenus {
		var menu model.Menu
		if err := db.Table("menu").Where("id = ?", roleMenu.MenuId).First(&menu).Error; err != nil {
			fmt.Println(err)
		}

		url := "/backend/" + strings.TrimPrefix(menu.Url, "/")
		if menu.Url != "/" && menu.Url != "" {
			if ok, err := enforcer.AddPolicy(PrefixRoleId+strconv.FormatInt(int64(roleId), 10), url, "GET|POST|PUT|DELETE"); err != nil {
				fmt.Println(ok, err.Error())
			}
		}
	}

}

// 检查用户是否拥有权限
func CheckUserPermission(userId, roleName, url, method string) (bool, error) {

	if roleName == "super_admin" {
		return true, nil
	}
	ok, err := Enforcer.Enforce(PrefixUserId+userId, url, method)
	return ok, err
}

func CheckRolePermission(roleId, roleName, url, method string) (bool, error) {

	if roleName == "super_admin" {
		return true, nil
	}

	ok, err := Enforcer.Enforce(PrefixRoleId+roleId, url, method)
	if err != nil {
		log.Fatal(err)
	}
	return ok, err
}
