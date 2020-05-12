package admin

import (
	"blog/model/auth"
	"blog/model/manage"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Row struct {
	Username	string
	RoleIds		string
}

type AdminRole struct {
	manage.Admin
	RoleIds string `json:"role_ids"`
}

func Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []AdminRole
	if err := db.Table("admin").Select("id,username").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}
	var rows []Row
	db.Table("admin").
		Select("admin.username, group_concat(admin_role.role_id) as role_ids").
		Joins("left join admin_role on admin_role.admin_id = admin.id").
		Group("admin.username").
		Find(&rows)
	fmt.Println("start")
	fmt.Println(rows)

	for k, v := range fields {
		for _, vv := range rows {
			if v.Username == vv.Username {
				fields[k].RoleIds = vv.RoleIds
			}
		}
	}

	helper.Success(ctx, fields)
}

func Show(ctx *gin.Context) {
	db := helper.Db()
	var field manage.Admin
	if err := db.Table("admin").Select("id,username").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {

	status := auth.Register(ctx, true)

	if !status {
		helper.Fail(ctx, "fail")
	}

	helper.Success(ctx, "success")
}

func Update(ctx *gin.Context) {
	db := helper.Db()
	var filed manage.Admin
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("admin").Model(&filed).Updates(requestJson).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field manage.Admin
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("admin").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}
