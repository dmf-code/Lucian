package adminRole

import (
	"blog/model/manage"
	"blog/utils/format"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func Index(ctx *gin.Context) {
	db := helper.Db()

	sql := "SELECT admin_role.id, role.name as role_name,role.id as role_id," +
		"admin.username as admin_name,admin.id as admin_id, admin_role.created_at, admin_role.updated_at FROM admin_role " +
		"LEFT JOIN role ON admin_role.role_id = role.id " +
		"LEFT JOIN admin ON admin_role.admin_id = admin.id;"

	fields, err := db.Raw(sql).Rows()

	if err != nil {
		fmt.Println(err)
		helper.Fail(ctx, err)
		return
	}
	var res []interface{}
	var roleIds []int

	for fields.Next() {
		var id,roleId,adminId int
		var roleName, adminName string
		var createdAt,updateAt format.JSONTime
		err = fields.Scan(&id, &roleName, &roleId, &adminName, &adminId, &createdAt, &updateAt)
		var item = struct{
			Id int `json:"id"`
			RoleName string `json:"role_name"`
			RoleId int `json:"role_id"`
			AdminName string `json:"admin_name"`
			AdminId int `json:"admin_id"`
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

	helper.Success(ctx, res, roleIds)
}

func Show(ctx *gin.Context) {
	db := helper.Db()
	var field manage.AdminRole
	if err := db.Table("admin_role").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {
	db := helper.Db()
	requestJson := helper.GetRequestJson(ctx)

	adminId := helper.Float64ToInt(requestJson["admin_id"].(float64))

	roleIdStr := requestJson["role_id"].(string)
	roleIds := strings.Split(roleIdStr, ",")

	sql := "INSERT INTO `admin_role` (`role_id`, `admin_id`, `created_at`) VALUES";
	for i, e := range roleIds {
		roleId, _ := strconv.Atoi(e)
		sql += fmt.Sprintf("('%d', '%d', '%s')", roleId, adminId, time.Now().Format("2006-01-02 15:04:05"))
		if i != len(roleIds) - 1 {
			sql += ","
		} else {
			sql += ";"
		}
	}
	fmt.Println(sql)
	if err := db.Exec(sql).Error; err != nil {
		fmt.Println(err)
		helper.Fail(ctx, "添加失败")
		return
	}

	helper.Success(ctx, "添加成功")
}

func Update(ctx *gin.Context) {
	db := helper.Db()
	var filed manage.AdminRole
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("admin_role").Model(&filed).Updates(requestJson).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field manage.AdminRole
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("admin_role").Unscoped().Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}
