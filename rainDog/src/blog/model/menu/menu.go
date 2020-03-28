package menu

import (
	"blog/model/manage"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Menu struct {

}

func Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []manage.Menu
	if err := db.Table("menu").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func Show(ctx *gin.Context) {
	var field manage.Menu
	if err := helper.Db().Table("menu").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {
	var field manage.Menu
	err := ctx.Bind(&field)
	fmt.Println(field)
	if err != nil {
		fmt.Println(err)
		helper.Fail(ctx, "绑定数据失败")
		return
	}
	if err = helper.Db().Table("menu").Create(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "success")
}

func Update(ctx *gin.Context) {
	var filed manage.Menu
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	if err := helper.Db().Table("menu").Model(&filed).Updates(requestJson).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	var field manage.Menu
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := helper.Db().Table("menu").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}

