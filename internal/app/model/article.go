package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rain/library/go-str"
	"rain/library/helper"
	"rain/library/response"
)

type Article struct {
	BaseModel
	MdCode      string `json:"mdCode" gorm:"type:longtext;column:md_code;comment:'markdown代码'"`
	HtmlCode    string `json:"htmlCode" gorm:"type:longtext;column:html_code;comment:'html代码'"`
	Title       string `json:"title" gorm:"column:title;comment:'标题'"`
	CategoryIds string `json:"categoryIds" gorm:"column:category_ids;comment:'分类id'"`
	TagIds      string `json:"tagIds" gorm:"column:tag_ids;comment:'标签id'"`
	Summary     string `json:"summary" gorm:"column:summary;comment:'简介'"`
	IsHide		int8	`json:"is_hide" gorm:"column:is_hide;default: 1;comment:'1.不隐藏 2.隐藏'"`
}

func (m *Article) Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Article
	page := ctx.Query("page")
	pageSize := ctx.Query("page_size")
	var total int
	db.Table("article").Where("deleted_at=null").Count(&total)
	if err := db.Table("article").Where("is_hide=1").Scopes(helper.Paginate(ctx)).Find(&fields).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}
	for k, v := range fields {
		fields[k].Summary = str.SubString(str.TrimHtml(v.HtmlCode), "...", 0, 120)
	}

	resp.Paginate(ctx, "ok", fields, total, str.ToInt(page), str.ToInt(pageSize))
}

func (m *Article) Show(ctx *gin.Context) {
	db := helper.Db()
	var field Article
	if err := db.Table("article").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", field)
}

func (m *Article) Store(ctx *gin.Context) {
	db := helper.Db()
	var field Article
	err := ctx.Bind(&field)
	fmt.Println(field)
	if err != nil {
		resp.Error(ctx, 400, "绑定数据失败")
		return
	}
	if err = db.Table("article").Create(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "ok")
}

func (m *Article) Update(ctx *gin.Context) {
	db := helper.Db()
	var filed Article
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = str.ToUint(ctx.Param("id"))
	if err := db.Table("article").Model(&filed).Updates(requestJson).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "ok")
}

func (m *Article) Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Article
	field.ID = str.ToUint(ctx.Param("id"))
	if err := db.Table("article").Delete(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx,"删除成功")
}
