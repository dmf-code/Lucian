package manage

import (
	"blog/model"
	"blog/utils/format"
	"github.com/jinzhu/gorm"
	"time"
)

type Menu struct {
	model.BaseModel
	Status		uint8	`gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"`					// 状态（1.启用 2.禁用）
	Memo		string	`gorm:"column:memo;size:64;" json:"memo" form:"memo"`											// 备注
	ParentID	uint64	`gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"`							// 父级ID
	Url			string	`gorm:"column:url;size:72;" json:"url" form:"url"`												// 菜单URL
	Name  		string	`gorm:"column:name;size:32;not null;" json:"name" form:"name"`									// 菜单名称
	Sequence	int		`gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`								// 排序值
	Type		uint8	`gorm:"column:type;type: tinyint(1);not null;" json:"type" form:"type"`							// 菜单类型
	Code		string	`gorm:"column:code;size:32;not null;unique_index:uk_menu_code;" json:"code" form:"code"`		// 菜单代码
	Icon		string	`gorm:"column:icon;size:32;" json:"icon" form:"icon"`											// icon
	OperateType	string	`gorm:"column:operate_type;size:32;not null;" json:"operate_type" form:"operate_type"`			// 操作类型 none/add/del/view/update
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
func (m *Menu) BeforeUpdate(scope gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

