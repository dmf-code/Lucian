package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rain/library/helper"
	"rain/library/response"
)

type Tag struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;comment:'tag name';"`
	Num  int    `json:"num"  gorm:"column:num;comment:'tag 使用次数';"`
}

func (m *Tag) Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Tag
	if err := db.Table("tag").Find(&fields).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", fields)
}

func (m *Tag) Show(ctx *gin.Context) {
	db := helper.Db()
	var field Tag
	if err := db.Table("tag").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", field)
}

func (m *Tag) Store(ctx *gin.Context) {
	db := helper.Db()
	var field Tag
	err := ctx.Bind(&field)
	field.Num = 0
	fmt.Println(field)
	if err != nil {
		resp.Error(ctx, 400, "绑定数据失败")
		return
	}
	if err = db.Table("tag").Create(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "success")
}

func (m *Tag) Update(ctx *gin.Context) {
	db := helper.Db()
	var filed Tag
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("tag").Model(&filed).Updates(requestJson).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "更新成功")
}

func (m *Tag) Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Tag
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("tag").Delete(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "删除成功")
}
