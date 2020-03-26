package permission

import (
	"blog/model/manage"
	"blog/utils/helper"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/jinzhu/gorm"
	"strconv"
)

const (
	PrefixUserId = "u"
	PrefixRoleId = "r"
)

var Enforcer *casbin.Enforcer

func Init() {
	enforcer, err := casbin.NewEnforcerSafe("conf/model.conf")
	if err != nil {
		fmt.Println(err)
	}

	var roles []manage.Role
	db := helper.Db("rain_dog")
	if err = db.Table("role").Find(&roles).Error; err != nil {
		fmt.Println(err)
	}

	if len(roles) == 0 {
		Enforcer = enforcer
	}

	for _, role := range roles {
		setRolePermission(db, enforcer, uint64(role.ID))
	}

	Enforcer = enforcer
}

// 设置角色权限
func setRolePermission(db *gorm.DB, enforcer *casbin.Enforcer, roleId uint64) {
	var roleMenus []manage.RoleMenu
	if err := db.Model(&manage.RoleMenu{RoleId: roleId}).Find(&roleMenus).Error; err != nil {
		fmt.Println(err)
	}
	for _, roleMenu := range roleMenus {
		var menu manage.Menu
		if err = db.Table("menu").Where("id = ?", roleMenu.MenuId).First(&menu).Error; err != nil {
			fmt.Println(err)
		}

		if menu.Type == 3 {
			enforcer.AddPermissionForUser(
				PrefixRoleId+strconv.FormatInt(int64(roleId), 10),
				"/backend"+menu.Url,
				"GET|POST|PUT|DELETE")
		}
	}
}
