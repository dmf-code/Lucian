package tutorial

import (
	"app/library/helper"
	"app/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Tutorial struct {
	model.BaseModel
	Img  		string		`json:"img" gorm:"size:256; column:img; comment:'教程封面图片'"`
	Title		string		`json:"title" gorm:"size:64; column:title; unique_index;comment: '标题';"`
	ParentId	int			`json:"parent_id" gorm:"column:parent_id; comment: '目录根节点';"`
	Type 		int			`json:"type" gorm:"column:type; comment: '类型：1.目录 2.菜单';"`
	Icon		string		`json:"icon" gorm:"column:icon;size:32;comment:'icon'" form:"icon"`
}

func Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Tutorial
	if err := db.Table("tutorial").Where("parent_id=?", 0).Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func Show(ctx *gin.Context) {
	db := helper.Db()
	var field Tutorial
	if err := db.Table("tutorial").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {
	db := helper.Db()
	var field Tutorial
	data := helper.GetRequestJson(ctx)
	fmt.Println(data)
	field.Type = helper.Float64ToInt(data["type"].(float64))
	field.Title = data["title"].(string)
	field.Icon = data["icon"].(string)
	field.ParentId = helper.Float64ToInt(data["parent_id"].(float64))
	fmt.Println(field)
	var content ContentTutorial
	content.HtmlCode = data["htmlCode"].(string)
	content.MdCode = data["mdCode"].(string)

	fmt.Println(11111)
	fmt.Println(content)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("tutorial").Create(&field).Error; err != nil {
			return err
		}

		if field.Type == 1 {
			return nil
		}

		if  err := tx.Table("content_tutorial").Create(&content).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		helper.Fail(ctx, err)
	}

	helper.Success(ctx, "success")
}

func Update(ctx *gin.Context) {
	db := helper.Db()
	var filed Tutorial
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	field.Type = helper.Float64ToInt(requestJson["type"].(float64))
	field.Title = requestJson["title"].(string)
	field.Icon = requestJson["icon"].(string)
	field.ParentId = helper.Float64ToInt(requestJson["parent_id"].(float64))

	if err := db.Table("tutorial").Model(&filed).Updates(requestJson).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Tutorial
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("tutorial").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}

