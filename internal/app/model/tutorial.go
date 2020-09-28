package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"rain/library/helper"
)

type ContentTutorial struct {
	BaseModel
	TutorialId int    `json:"tutorial_id" gorm:"type:int;column:tutorial_id;comment:'菜单id'"`
	MdCode     string `json:"mdCode" gorm:"type:longtext;column:md_code;comment:'markdown代码'"`
	HtmlCode   string `json:"htmlCode" gorm:"type:longtext;column:html_code;comment:'html代码'"`
}


func (m *ContentTutorial) getMenuTree(pid int) []*TreeList {
	db := helper.Db()
	var menus []Tutorial
	if err := db.Preload("ContentTutorial").Where("parent_id = ?", pid).Find(&menus).Error; err != nil {
		panic(err)
	}
	var treeList []*TreeList
	for _, v := range menus {
		// 筛除非选中节点
		child := m.getMenuTree(int(v.ID))
		node := &TreeList{
			Id:       v.ID,
			Name:     v.Title,
			Value:    uint64(v.ID),
			Label:    v.Title,
			Type:     uint8(v.Type),
			Pid:      uint64(v.ParentId),
			Icon:     v.Icon,
			MdCode:   v.ContentTutorial.MdCode,
			HtmlCode: v.ContentTutorial.HtmlCode,
		}

		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}

func (m *ContentTutorial) List(ctx *gin.Context) {
	t := ctx.Param("pid")

	treeList := m.getMenuTree(helper.Str2Int(t))
	helper.Success(ctx, treeList)
}

type Tutorial struct {
	BaseModel
	Img             string `json:"img" gorm:"size:256; column:img; comment:'教程封面图片'"`
	Title           string `json:"title" gorm:"size:64; column:title; unique_index;comment: '标题';"`
	ParentId        int    `json:"parent_id" gorm:"column:parent_id; comment: '目录根节点';"`
	Type            int    `json:"type" gorm:"column:type; comment: '类型：1.目录 2.菜单';"`
	Icon            string `json:"icon" gorm:"column:icon;size:32;comment:'icon'" form:"icon"`
	ContentTutorial ContentTutorial
}

func (m *Tutorial) Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Tutorial
	if err := db.Table("tutorial").Where("parent_id=?", 0).Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func (m *Tutorial) Show(ctx *gin.Context) {
	db := helper.Db()
	var field Tutorial
	if err := db.Table("tutorial").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func (m *Tutorial) Store(ctx *gin.Context) {
	db := helper.Db()
	var field Tutorial
	data := helper.GetRequestJson(ctx)

	field.Type = helper.Float64ToInt(data["type"].(float64))
	field.Title = data["title"].(string)
	field.Icon = data["icon"].(string)
	field.ParentId = helper.Float64ToInt(data["parent_id"].(float64))
	if field.ParentId == 0 {
		field.Img = data["img"].(string)
	}

	var content ContentTutorial

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("tutorial").Create(&field).Error; err != nil {
			return err
		}

		if field.Type == 1 {
			return nil
		}
		content.TutorialId = int(field.ID)
		content.HtmlCode = data["htmlCode"].(string)
		content.MdCode = data["mdCode"].(string)
		if err := tx.Table("content_tutorial").Create(&content).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		helper.Fail(ctx, err)
	}

	helper.Success(ctx, "success")
}

func (m *Tutorial) Update(ctx *gin.Context) {
	db := helper.Db()

	requestJson := helper.GetRequestJson(ctx)
	var field Tutorial
	field.ID = helper.Str2Uint(ctx.Param("id"))
	field.Type = helper.Float64ToInt(requestJson["type"].(float64))
	field.Title = requestJson["title"].(string)
	field.Icon = requestJson["icon"].(string)
	field.ParentId = helper.Float64ToInt(requestJson["parent_id"].(float64))
	if field.ParentId == 0 {
		field.Img = requestJson["img"].(string)
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("tutorial").Model(&field).Updates(Tutorial{
			ParentId: field.ParentId,
			Title:    field.Title,
			Icon:     field.Icon,
			Type:     field.Type,
			Img:      field.Img,
		}).Error; err != nil {
			return err
		}

		if field.ParentId == 0 || field.Type == 1 {
			return nil
		}

		var content ContentTutorial
		content.HtmlCode = requestJson["htmlCode"].(string)
		content.MdCode = requestJson["mdCode"].(string)

		if err := tx.Table("content_tutorial").
			Model(&content).
			Where("tutorial_id = ?", field.ID).
			Updates(ContentTutorial{
				MdCode:   content.MdCode,
				HtmlCode: content.HtmlCode,
			}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helper.Fail(ctx, err)
		return
	}

	helper.Success(ctx, "更新成功")
}

func (m *Tutorial) Destroy(ctx *gin.Context) {
	db := helper.Db()
	var field Tutorial
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := db.Table("tutorial").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}

