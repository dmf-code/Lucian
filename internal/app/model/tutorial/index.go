package tutorial

import (
	"app/library/helper"
	"app/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CoverTutorial struct {
	model.BaseModel
	Img  		string		`json:"img" gorm:"size:256; column:img; comment:'教程封面图片'"`
	Title		string		`json:"title" gorm:"size:64; column:title; unique_index;comment: '标题';"`
	Root		int			`json:"root" gorm:"column:root; comment: '目录根节点';"`
}

type CoverMenu struct {
	Id		int 		`json:"id"`
	Label	string		`json:"label"`
}

func Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []CoverTutorial
	if err := db.Table("cover_tutorial").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func Show(ctx *gin.Context) {
	db := helper.Db()
	var field CoverTutorial
	if err := db.Table("cover_tutorial").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {
	db := helper.Db()
	var field CoverTutorial
	err := ctx.Bind(&field)
	fmt.Println(field)
	if err != nil {
		helper.Fail(ctx, "绑定数据失败")
		return
	}
	if err = db.Table("cover_tutorial").Create(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "success")
}

func Update(ctx *gin.Context) {
	db := helper.Db()
	var filed CoverTutorial
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("cover_tutorial").Model(&filed).Updates(requestJson).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field CoverTutorial
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("cover_tutorial").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}

func CoverMenuList(ctx *gin.Context)  {
	db := helper.Db()
	var fields []CoverTutorial
	if err := db.Table("cover_tutorial").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}
	var menus []*CoverMenu
	for _, v := range fields {
		menu := &CoverMenu{
			Id:    int(v.ID),
			Label: v.Title,
		}
		menus = append(menus, menu)
	}

	helper.Success(ctx, menus)
}
