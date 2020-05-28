package article

import (
	"app/model"
	"app/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Article struct {
	model.BaseModel
	MdCode		string `json:"mdCode" gorm:"column:md_code;comment:'markdown代码'"`
	HtmlCode	string `json:"htmlCode" gorm:"column:html_code;comment:'html代码'"`
	Title		string `json:"title" gorm:"column:title;comment:'标题'"`
	CategoryIds string `json:"categoryIds" gorm:"column:category_ids;comment:'分类id'"`
	TagIds		string `json:"tagIds" gorm:"column:tag_ids;comment:'标签id'"`
}


func Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Article
	if err := db.Table("article").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func Show(ctx *gin.Context) {
	db := helper.Db()
	var field Article
	if err := db.Table("article").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {
	db := helper.Db()
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
	db := helper.Db()
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
	db := helper.Db()
	var field Article
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("article").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}
