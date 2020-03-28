package tag

import (
	"blog/model"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Tag struct {
	model.BaseModel
	Name string `json:"name" gorm:"column:name;"`
	Num int `json:"num"  gorm:"column:num;"`
}


func Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Tag
	if err := db.Table("tag").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func Show(ctx *gin.Context) {
	db := helper.Db()
	var field Tag
	if err := db.Table("tag").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {
	db := helper.Db()
	var field Tag
	err := ctx.Bind(&field)
	field.Num = 0
	fmt.Println(field)
	if err != nil {
		helper.Fail(ctx, "绑定数据失败")
		return
	}
	if err = db.Table("tag").Create(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "success")
}

func Update(ctx *gin.Context) {
	db := helper.Db()
	var filed Tag
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("tag").Model(&filed).Updates(requestJson).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Tag
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("tag").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}


