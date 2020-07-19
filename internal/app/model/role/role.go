package role

import (
	"app/bootstrap/Table"
	"app/library/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
)

func Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Table.Role
	if err := db.Table("role").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func Show(ctx *gin.Context) {
	db := helper.Db()
	var field Table.Role
	if err := db.Table("role").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {
	db := helper.Db()
	var field Table.Role
	requestJson := helper.GetRequestJson(ctx)
	field.Memo = requestJson["memo"].(string)
	field.Name = requestJson["name"].(string)
	menuIds := strings.Split(requestJson["menus"].(string), ",")
	roleId, _ := ctx.Get("roleId")
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("role").Create(&field).Error; err != nil {
			return err
		}
		for menuId, _ := range menuIds {
			err := tx.Table("role_menu").Create(&Table.RoleMenu{
				RoleId: uint64(roleId.(int)),
				MenuId: uint64(menuId),
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		helper.Fail(ctx, "failed")
		return
	}

	helper.Success(ctx, "success")
}

func Update(ctx *gin.Context) {
	db := helper.Db()
	var filed Table.Role
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	filed.Name = requestJson["name"].(string)
	filed.Memo = requestJson["memo"].(string)
	menuIds := strings.Split(requestJson["menus"].(string), ",")
	roleId, _ := ctx.Get("roleId")
	fmt.Println(roleId)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("role").Model(&filed).Update("name", "memo").Error; err != nil {
			helper.Fail(ctx, err.Error())
			return err
		}
		tx.Where("role_id = ?", filed.ID).Delete(&Table.RoleMenu{})
		for _, menuId := range menuIds {
			err := tx.Table("role_menu").Create(&Table.RoleMenu{
				RoleId: uint64(filed.ID),
				MenuId: uint64(helper.Str2Uint(menuId)),
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		helper.Fail(ctx, "failed")
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Table.Role
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("role").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}
