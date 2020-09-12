package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"rain/library/format"
	"rain/library/helper"
	"strings"
	"time"
)

type Role struct {
	BaseModel
	Name     string `gorm:"column:name;size:32;not null;comment:'名称';unique_index;" json:"name" form:"name"`
	Memo     string `gorm:"column:memo;size:64;comment:'备注'" json:"memo" form:"memo"`
	Sequence uint64 `gorm:"column:sequence;not null;default: 5;comment:'排序值'" json:"sequence" form:"sequence"`
	//ParentId	uint64	`gorm:"column:parent_id;not null;comment:'父级ID'" json:"parent_id" form:"parent_id"`
}

func (Role) TableName() string {
	return TableName("role")
}

// 添加之前
func (m *Role) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新之前
func (m *Role) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

func (m *Role) Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Role
	if err := db.Table("role").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func (m *Role) Show(ctx *gin.Context) {
	db := helper.Db()
	var field Role
	if err := db.Table("role").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func (m *Role) Store(ctx *gin.Context) {
	db := helper.Db()
	var field Role
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
			err := tx.Table("role_menu").Create(&RoleMenu{
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

func (m *Role) Update(ctx *gin.Context) {
	db := helper.Db()
	var filed Role
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
		tx.Table("role_menu").Where("role_id = ?", filed.ID).Delete(&RoleMenu{})
		for _, menuId := range menuIds {
			err := tx.Table("role_menu").Create(&RoleMenu{
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

func (m *Role) Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Role
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("role").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}
