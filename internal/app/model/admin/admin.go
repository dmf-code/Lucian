package admin

import (
	"app/bootstrap/Table"
	"app/model/auth"
	"app/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
)

type Row struct {
	Username	string
	RoleIds		string
}

type AdminRole struct {
	Table.Admin
	RoleIds string `json:"role_ids"`
}
//// 添加之前
//func (m *AdminRole) BeforeCreate(scope *gorm.Scope) error {
//	m.CreatedAt = format.JSONTime{Time: time.Now()}
//	m.UpdatedAt = format.JSONTime{Time: time.Now()}
//	return nil
//}
//
//// 更新之前
//func (m *AdminRole) BeforeUpdate(scope *gorm.Scope) error {
//	m.UpdatedAt = format.JSONTime{Time: time.Now()}
//	return nil
//}

func Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []AdminRole
	if err := db.Table("admin").Select("id,username,created_at,updated_at").Find(&fields).Error; err != nil {
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
	var field Table.Admin
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

func updateAdmin(filed Table.Admin, requestJson map[string]interface{}) error {
	db := helper.Db()
	return db.Transaction(func(tx *gorm.DB) error {
		delete(requestJson, "password")
		delete(requestJson, "createdAt")
		delete(requestJson, "updatedAt")
		if err := tx.Table("admin").Model(&filed).Updates(requestJson).Error; err != nil {
			return err
		}
		if err := tx.Table("admin_role").Where("admin_id = ?", filed.ID).Delete(Table.AdminRole{}).Error; err != nil {
			return err
		}

		roleIds := strings.Split(requestJson["role_ids"].(string), ",")
		fmt.Println(roleIds)
		for _, v := range roleIds {
			if err := tx.Table("admin_role").Create(&Table.AdminRole{AdminId: uint64(filed.ID), RoleId: uint64(helper.Str2Uint(v))}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func Update(ctx *gin.Context) {
	var filed Table.Admin
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))

	if err := updateAdmin(filed, requestJson); err != nil {
		fmt.Println(err)
		helper.Fail(ctx, "更新失败")
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Table.Admin
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("admin").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}
