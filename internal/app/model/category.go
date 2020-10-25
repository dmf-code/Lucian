package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rain/library/helper"
	"rain/library/response"
)

type Category struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;comment: '分类名';"`
	Num  int    `json:"num"  gorm:"column:num;comment:'分类使用次数';"`
}

func (m *Category) Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Category
	if err := db.Table("category").Find(&fields).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", fields)
}

func (m *Category) Show(ctx *gin.Context) {
	db := helper.Db()
	var field Category
	if err := db.Table("category").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", field)
}

func (m *Category) Store(ctx *gin.Context) {
	db := helper.Db()
	var field Category
	err := ctx.Bind(&field)
	field.Num = 0
	fmt.Println(field)
	if err != nil {
		resp.Error(ctx, 400, "绑定数据失败")
		return
	}
	if err = db.Table("category").Create(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "ok")
}

func (m *Category) Update(ctx *gin.Context) {
	db := helper.Db()
	var filed Category
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("category").Model(&filed).Updates(requestJson).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "更新成功")
}

func (m *Category) Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Category
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("category").Delete(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx,"删除成功")
}
