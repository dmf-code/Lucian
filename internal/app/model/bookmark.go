package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"rain/library/format"
	"rain/library/go-str"
	"rain/library/helper"
	resp "rain/library/response"
	"time"
)

type Bookmark struct{
	BaseModel
	Name		string		`gorm:"comment: '名称'" json:"name"`
	Url 		string		`gorm:"comment: '书签url'" json:"url"`
	IsHide		int			`gorm:"default:1; comment: '1.不隐藏 2.隐藏'" json:"is_hide"`
}


func (Bookmark) TableName() string {
	return TableName("Bookmark")
}

// 添加之前
func (m *Bookmark) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新之前
func (m *Bookmark) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

func (m *Bookmark) Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Bookmark
	Bookmark := db.Table("bookmark")
	isHide := ctx.DefaultQuery("is_hide", "none")

	if isHide != "none" {
		Bookmark = Bookmark.Where("is_hide = ?", isHide)
	}

	if err := Bookmark.Find(&fields).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", fields)
}

func (m *Bookmark) Show(ctx *gin.Context) {
	db := helper.Db()
	var field Bookmark
	if err := db.Table("bookmark").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", field)
}

func (m *Bookmark) Store(ctx *gin.Context) {
	db := helper.Db()
	var field Bookmark
	requestJson := helper.GetRequestJson(ctx)
	field.Url = requestJson["url"].(string)
	field.Name = requestJson["name"].(string)
	field.IsHide = helper.Float64ToInt(requestJson["is_hide"].(float64))
	if err := db.Table("bookmark").Create(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "ok")
}

func (m *Bookmark) Update(ctx *gin.Context) {
	db := helper.Db()
	var field Bookmark
	requestJson := helper.GetRequestJson(ctx)
	field.ID = str.ToUint(ctx.Param("id"))
	field.Name = requestJson["name"].(string)
	field.Url = requestJson["url"].(string)
	field.IsHide = helper.Float64ToInt(requestJson["is_hide"].(float64))

	if err := db.Table("bookmark").Model(&field).Updates(
		map[string]interface{}{
			"name": field.Name,
			"url": field.Url,
			"is_hide": field.IsHide}).Error; err != nil {
			resp.Error(ctx, 400, err.Error())
			return
	}

	resp.Success(ctx, "更新成功")
}

func (m *Bookmark) Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Bookmark
	field.ID = str.ToUint(ctx.Param("id"))
	if err := db.Table("bookmark").Delete(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "删除成功")
}
