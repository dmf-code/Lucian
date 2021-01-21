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

type Nav struct{
	BaseModel
	Name		string		`gorm:"comment: '名称'" json:"name"`
	Path 		string		`gorm:"comment: '路由path'" json:"path"`
	IsHide		int			`gorm:"default:1; comment: '1.不隐藏 2.隐藏'" json:"is_hide"`
}


func (Nav) TableName() string {
	return TableName("Nav")
}

// 添加之前
func (m *Nav) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新之前
func (m *Nav) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

func (m *Nav) Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Nav
	nav := db.Table("Nav")
	isHide := ctx.DefaultQuery("is_hide", "none")

	if isHide != "none" {
		nav = nav.Where("is_hide = ?", isHide)
	}

	if err := nav.Find(&fields).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", fields)
}

func (m *Nav) Show(ctx *gin.Context) {
	db := helper.Db()
	var field Nav
	if err := db.Table("Nav").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", field)
}

func (m *Nav) Store(ctx *gin.Context) {
	db := helper.Db()
	var field Nav
	requestJson := helper.GetRequestJson(ctx)
	field.Path = requestJson["path"].(string)
	field.Name = requestJson["name"].(string)
	field.IsHide = helper.Float64ToInt(requestJson["is_hide"].(float64))
	if err := db.Table("Nav").Create(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "ok")
}

func (m *Nav) Update(ctx *gin.Context) {
	db := helper.Db()
	var field Nav
	requestJson := helper.GetRequestJson(ctx)
	field.ID = str.ToUint(ctx.Param("id"))
	field.Name = requestJson["name"].(string)
	field.Path = requestJson["path"].(string)
	field.IsHide = helper.Float64ToInt(requestJson["is_hide"].(float64))

	if err := db.Table("Nav").Model(&field).Updates(
		map[string]interface{}{
			"name": field.Name,
			"path": field.Path,
			"is_hide": field.IsHide}).Error; err != nil {
			resp.Error(ctx, 400, err.Error())
			return
	}

	resp.Success(ctx, "更新成功")
}

func (m *Nav) Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Nav
	field.ID = str.ToUint(ctx.Param("id"))
	if err := db.Table("Nav").Delete(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "删除成功")
}
