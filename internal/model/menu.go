package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"rain/library/format"
	"rain/library/go-str"
	"rain/library/helper"
	"rain/library/response"
	"time"
)

type TreeList struct {
	Id        uint        `json:"id"`
	Name      string      `json:"name"`
	Pid       uint64      `json:"pid"`
	Label     string      `json:"label"` //冗余前端字段
	Value     uint64      `json:"value"` //冗余前端字段
	Status    int         `json:"status"`
	Type      uint8       `json:"type"`
	Memo      string      `json:"memo"`
	Sequence  int         `json:"sequence"`
	Url       string      `json:"url"`
	FullUrl   string      `json:"full_url"`
	Component string      `json:"component"`
	Icon      string      `json:"icon"`
	MdCode     string 	  `json:"mdCode" gorm:"type:longtext;column:md_code;comment:'markdown代码'"`
	HtmlCode   string 	  `json:"htmlCode" gorm:"type:longtext;column:html_code;comment:'html代码'"`
	OperateType string 	  `gorm:"column:operate_type;size:32;not null;comment:'操作类型 none/add/del/view/update'" json:"operate_type" form:"operate_type"`
	Children  []*TreeList `json:"children,omitempty"`
}

type Menu struct {
	BaseModel
	Status      uint8  `gorm:"column:status;type:tinyint(1);not null;comment:'状态（1.启用 2.禁用）'" json:"status" form:"status"`
	Memo        string `gorm:"column:memo;size:64;comment:'备注'" json:"memo" form:"memo"`
	ParentID    uint64 `gorm:"column:parent_id;not null;comment:'父级ID'" json:"parent_id" form:"parent_id"`
	Url         string `gorm:"column:url;size:72;comment:'菜单URL'" json:"url" form:"url"`
	Name        string `gorm:"column:name;size:32;not null;comment:'菜单名称'" json:"name" form:"name"`
	Sequence    int    `gorm:"column:sequence;not null;comment:'排序值'" json:"sequence" form:"sequence"`
	Type        uint8  `gorm:"column:type;type: tinyint(1);not null;comment:'菜单类型 (1.目录 2.菜单 3.按钮 4.接口)'" json:"type" form:"type"`
	Component   string `gorm:"column:component;size:255;not null;comment:'组件路径'" json:"component" form:"component"`
	Icon        string `gorm:"column:icon;size:32;comment:'icon'" json:"icon" form:"icon"`
	OperateType string `gorm:"column:operate_type;size:32;not null;comment:'操作类型 none/add/del/view/update'" json:"operate_type" form:"operate_type"`
}

func (Menu) TableName() string {
	return TableName("menu")
}

// 添加之前
func (m *Menu) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新之前
func (m *Menu) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}


func (m *Menu) getMenu(pid int, path string) []*TreeList {
	var menus []Menu
	if pid == 0 {
		helper.Db().
			Table("menu").
			Where("parent_id = ?", pid).
			Where("url = ?", "/admin").
			Order("sequence").
			Find(&menus)
	} else {
		helper.Db().Table("menu").Where("parent_id = ?", pid).Order("sequence").Find(&menus)
	}
	var treeList []*TreeList
	for _, v := range menus {
		if v.Type >= 3 {
			continue
		}
		var fullPath string
		if path != "" {
			fullPath = path + "/" + v.Url
		} else {
			fullPath = v.Url
		}

		child := m.getMenu(int(v.ID), fullPath)

		node := &TreeList{
			Id:        		v.ID,
			Name:      		v.Name,
			Label:     		v.Name,
			Value:     		uint64(v.ID),
			Status:    		int(v.Status),
			Type:      		v.Type,
			Memo:      		v.Memo,
			Component: 		v.Component,
			Sequence:  		v.Sequence,
			Url:       		v.Url,
			FullUrl:   		fullPath,
			Pid:       		v.ParentID,
			Icon:      		v.Icon,
			OperateType: 	v.OperateType,
		}

		node.Children = child
		treeList = append(treeList, node)

	}

	return treeList
}

func (m *Menu) getApi(pid int, path string) []*TreeList {
	var menus []Menu
	if pid == 0 {
		helper.Db().
			Table("menu").
			Where("parent_id = ?", pid).
			Order("sequence").
			Find(&menus)
	} else {
		helper.Db().Table("menu").Where("parent_id = ?", pid).Order("sequence").Find(&menus)
	}
	var treeList []*TreeList
	for _, v := range menus {

		var fullPath string
		if path != "" {
			fullPath = path + "/" + v.Url
		} else {
			fullPath = v.Url
		}

		child := m.getMenu(int(v.ID), fullPath)
		node := &TreeList{
			Id:        v.ID,
			Name:      v.Name,
			Label:     v.Name,
			Value:     uint64(v.ID),
			Component: v.Component,
			Sequence:  v.Sequence,
			Url:       v.Url,
			FullUrl:   fullPath,
			Pid:       v.ParentID,
			Icon:      v.Icon,
			OperateType: 	v.OperateType,
		}

		node.Children = child
		treeList = append(treeList, node)

	}

	return treeList
}

func (m *Menu) Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Menu
	if err := db.Table("menu").Find(&fields).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx,"ok", fields)
}

func (m *Menu) Show(ctx *gin.Context) {
	var field Menu
	if err := helper.Db().Table("menu").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		resp.Error(ctx, 400, "查询失败")
		return
	}

	resp.Success(ctx, "ok", field)
}

func (m *Menu) Store(ctx *gin.Context) {
	var field Menu
	err := ctx.Bind(&field)
	fmt.Println(field)
	if err != nil {
		fmt.Println(err)
		resp.Error(ctx, 400, "绑定数据失败")
		return
	}
	if err = helper.Db().Table("menu").Create(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "ok")
}

func (m *Menu) Update(ctx *gin.Context) {
	var filed Menu
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = str.ToUint(ctx.Param("id"))
	fmt.Println(requestJson)
	if err := helper.Db().Table("menu").Model(&filed).Updates(requestJson).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx, "更新成功")
}

func (m *Menu) Destroy(ctx *gin.Context) {
	var field Menu
	field.ID = str.ToUint(ctx.Param("id"))
	if err := helper.Db().Table("menu").Delete(&field).Error; err != nil {
		resp.Error(ctx, 400, err.Error())
		return
	}

	resp.Success(ctx,"删除成功")
}

func (m *Menu) List(ctx *gin.Context) {
	treeList := m.getMenu(0, "")
	resp.Success(ctx,  "ok", treeList)
}

func (m *Menu) ApiList(ctx *gin.Context) {
	treeList := m.getApi(0, "")
	resp.Success(ctx,  "ok", treeList)
}
