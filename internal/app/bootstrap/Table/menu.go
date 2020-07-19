package Table

import (
	"app/model"
	"app/library/format"
	"github.com/jinzhu/gorm"
	"time"
)

type Menu struct {
	model.BaseModel
	Status		uint8	`gorm:"column:status;type:tinyint(1);not null;comment:'状态（1.启用 2.禁用）'" json:"status" form:"status"`
	Memo		string	`gorm:"column:memo;size:64;comment:'备注'" json:"memo" form:"memo"`
	ParentID	uint64	`gorm:"column:parent_id;not null;comment:'父级ID'" json:"parent_id" form:"parent_id"`
	Url			string	`gorm:"column:url;size:72;comment:'菜单URL'" json:"url" form:"url"`
	Name  		string	`gorm:"column:name;size:32;not null;comment:'菜单名称'" json:"name" form:"name"`
	Sequence	int		`gorm:"column:sequence;not null;comment:'排序值'" json:"sequence" form:"sequence"`
	Type		uint8	`gorm:"column:type;type: tinyint(1);not null;comment:'菜单类型 (1.目录 2.菜单 3.按钮 4.接口)'" json:"type" form:"type"`
	Component	string	`gorm:"column:component;size:255;not null;comment:'组件路径'" json:"component" form:"component"`
	Icon		string	`gorm:"column:icon;size:32;comment:'icon'" json:"icon" form:"icon"`
	OperateType	string	`gorm:"column:operate_type;size:32;not null;comment:'操作类型 none/add/del/view/update'" json:"operate_type" form:"operate_type"`
}

func (Menu) TableName()string {
	return model.TableName("menu")
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

