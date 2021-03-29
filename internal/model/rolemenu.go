package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"rain/library/format"
	"rain/library/go-str"
	"rain/library/helper"
	"rain/library/response"
	"strconv"
	"strings"
	"time"
)

type RoleMenu struct {
	BaseModel
	RoleId uint64 `gorm:"column:role_id;unique_index:uk_role_menu_role_id;not null;comment:'角色ID'"`
	MenuId uint64 `gorm:"column:menu_id;unique_index:uk_role_menu_role_id;not null;comment:'菜单ID'"`
}

func (RoleMenu) TableName() string {
	return TableName("role_menu")
}

// 添加前
func (m *RoleMenu) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新前
func (m *RoleMenu) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

func (m *RoleMenu) Index(ctx *gin.Context) {
	db := helper.Db()

	sql := "SELECT role_menu.id, role.name as role_name,role.id as role_id," +
		"menu.name as menu_name,menu.id as menu_id, role_menu.created_at, role_menu.updated_at FROM role_menu " +
		"LEFT JOIN role ON role_menu.role_id = role.id " +
		"LEFT JOIN menu ON role_menu.menu_id = menu.id;"

	fields, err := db.Raw(sql).Rows()

	if err != nil {
		fmt.Println(err)
		resp.Error(ctx, 400, err.Error())
		return
	}
	var res []interface{}
	var menuIds []int

	for fields.Next() {
		var id, roleId, menuId int
		var roleName, menuName string
		var createdAt, updateAt format.JSONTime
		err = fields.Scan(&id, &roleName, &roleId, &menuName, &menuId, &createdAt, &updateAt)
		var item = struct {
			Id        int    `json:"id"`
			RoleName  string `json:"role_name"`
			RoleId    int    `json:"role_id"`
			MenuName  string `json:"menu_name"`
			MenuId    int    `json:"menu_id"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at,omitempty"`
		}{
			id,
			roleName,
			roleId,
			menuName,
			menuId,
			createdAt.Format(format.TimeFormat),
			updateAt.Format(format.TimeFormat),
		}

		menuIds = append(menuIds, menuId)
		fmt.Println(item)
		res = append(res, item)
		fmt.Println(err)
	}

	resp.Success(ctx, "ok", res, menuIds)
}

func (m *RoleMenu) Show(ctx *gin.Context) {
	db := helper.Db()
	var field RoleMenu
	if err := db.Table("role_menu").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", field)
}

func (m *RoleMenu) Store(ctx *gin.Context) {
	db := helper.Db()
	requestJson := helper.GetRequestJson(ctx)

	roleId := helper.Float64ToInt(requestJson["role_id"].(float64))

	menuIdStr := requestJson["menu_id"].(string)
	menuId := strings.Split(menuIdStr, ",")

	sql := "INSERT INTO `role_menu` (`role_id`, `menu_id`, `created_at`) VALUES"
	var mId int
	for i, e := range menuId {
		mId, _ = strconv.Atoi(e)
		sql += fmt.Sprintf("('%d', '%d', '%s')", roleId, mId, time.Now().Format("2006-01-02 15:04:05"))
		if i != len(menuId)-1 {
			sql += ","
		} else {
			sql += ";"
		}
	}
	fmt.Println(sql)
	if err := db.Exec(sql).Error; err != nil {
		fmt.Println(err)
		resp.Error(ctx, 400, "添加失败")
		return
	}

	resp.Success(ctx, "添加成功")
}

func (m *RoleMenu) Update(ctx *gin.Context) {
	db := helper.Db()
	var filed RoleMenu
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = str.ToUint(ctx.Param("id"))
	if err := db.Table("role_menu").Model(&filed).Updates(requestJson).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "更新成功")
}

func (m *RoleMenu) Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field RoleMenu
	field.ID = str.ToUint(ctx.Param("id"))
	if err := db.Table("role_menu").Unscoped().Delete(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "删除成功")
}

func (m *RoleMenu) List(ctx *gin.Context) {
	db := helper.Db()
	var fields []RoleMenu
	if roleId := ctx.DefaultQuery("roleId", ""); roleId != "" {
		db = db.Table("role_menu").Where("role_id = ?", roleId)
		fmt.Println(roleId)
	}
	db.Find(&fields)
	fmt.Println(fields)
	resp.Success(ctx, "ok", fields)
}
