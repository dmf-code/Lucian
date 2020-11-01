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

type AdminRole struct {
	BaseModel
	AdminId uint64 `gorm:"column:admin_id;unique_index:uk_admin_role_admin_id;not null;comment='管理员ID'"`
	RoleId  uint64 `gorm:"column:role_id;unique_index:uk_admin_role_admin_id;not null;comment='角色ID'"`
}

// 表名
func (AdminRole) TableName() string {
	return TableName("admin_role")
}

// 添加前
func (m *AdminRole) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新前
func (m *AdminRole) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

func (m *AdminRole) Index(ctx *gin.Context) {
	db := helper.Db()

	sql := "SELECT admin_role.id, role.name as role_name,role.id as role_id," +
		"admin.username as admin_name,admin.id as admin_id, admin_role.created_at, admin_role.updated_at FROM admin_role " +
		"LEFT JOIN role ON admin_role.role_id = role.id " +
		"LEFT JOIN admin ON admin_role.admin_id = admin.id;"

	fields, err := db.Raw(sql).Rows()

	if err != nil {
		fmt.Println(err)
		resp.Error(ctx, 400, err.Error())
		return
	}
	var res []interface{}
	var roleIds []int

	for fields.Next() {
		var id, roleId, adminId int
		var roleName, adminName string
		var createdAt, updateAt format.JSONTime
		err = fields.Scan(&id, &roleName, &roleId, &adminName, &adminId, &createdAt, &updateAt)
		var item = struct {
			Id        int    `json:"id"`
			RoleName  string `json:"role_name"`
			RoleId    int    `json:"role_id"`
			AdminName string `json:"admin_name"`
			AdminId   int    `json:"admin_id"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at,omitempty"`
		}{
			id,
			roleName,
			roleId,
			adminName,
			adminId,
			createdAt.Format(format.TimeFormat),
			updateAt.Format(format.TimeFormat),
		}

		roleIds = append(roleIds, roleId)
		fmt.Println(item)
		res = append(res, item)
		fmt.Println(err)
	}

	resp.Success(ctx, "ok", res, roleIds)
}

func (m *AdminRole) Show(ctx *gin.Context) {
	db := helper.Db()
	var fields []int
	if err := db.Table("admin_role").Where("admin_id = ?", ctx.Param("id")).Pluck("role_id", &fields).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", fields)
}

func (m *AdminRole) Store(ctx *gin.Context) {
	db := helper.Db()
	requestJson := helper.GetRequestJson(ctx)

	adminId := helper.Float64ToInt(requestJson["admin_id"].(float64))

	db.Unscoped().Where("admin_id = ?", adminId).Delete(&AdminRole{})

	roleIdStr := requestJson["role_id"].(string)
	roleIds := strings.Split(roleIdStr, ",")

	sql := "INSERT INTO `admin_role` (`role_id`, `admin_id`, `created_at`) VALUES"
	for i, e := range roleIds {
		roleId, _ := strconv.Atoi(e)
		sql += fmt.Sprintf("('%d', '%d', '%s')", roleId, adminId, time.Now().Format("2006-01-02 15:04:05"))
		if i != len(roleIds)-1 {
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

func (m *AdminRole) Update(ctx *gin.Context) {
	db := helper.Db()
	var filed AdminRole
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = str.ToUint(ctx.Param("id"))
	if err := db.Table("admin_role").Model(&filed).Updates(requestJson).Error; err != nil {
		resp.Error(ctx,400, err.Error())
		return
	}

	resp.Success(ctx, "更新成功")
}

func (m *AdminRole) Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field AdminRole
	field.ID = str.ToUint(ctx.Param("id"))
	if err := db.Table("admin_role").Unscoped().Delete(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx,  "删除成功")
}
