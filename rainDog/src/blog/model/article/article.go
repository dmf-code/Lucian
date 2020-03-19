package article

import (
	"blog/utils/helper"
	"blog/utils/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Article struct {
	model.BaseModel
	MdCode		string `json:"mdCode" gorm:"column:md_code;"`
	HtmlCode	string `json:"htmlCode" gorm:"html_code;"`
	Title		string `json:"title"`
	CategoryIds string `json:"categoryIds" gorm:"column:category_ids;"`
	TagIds		string `json:"tagIds" gorm:"column:tag_ids;"`
}


func Index(ctx *gin.Context) {
	db := helper.Db("rain_dog")
	var fields []Article
	if err := db.Table("article").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func Show(ctx *gin.Context) {
	db := helper.Db("rain_dog")
	var field Article
	if err := db.Table("article").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {
	db := helper.Db("rain_dog")
	var field Article
	err := ctx.Bind(&field)
	fmt.Println(field)
	if err != nil {
		helper.Fail(ctx, "绑定数据失败")
		return
	}
	if err = db.Table("article").Create(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "success")
}

func Update(ctx *gin.Context) {
	db := helper.Db("rain_dog")
	var filed Article
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("article").Model(&filed).Updates(requestJson).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	db := helper.Db("rain_dog")
	var field Article
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("article").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}
