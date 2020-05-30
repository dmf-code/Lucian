package Table

import (
	"app/model"
	"app/utils/format"
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	model.BaseModel
	Name		string	`gorm:"column:name;size:32;not null;comment:'名称'" json:"name" form:"name"`
	Memo 		string	`gorm:"column:memo;size:64;comment:'备注'" json:"memo" form:"memo"`
	Sequence	uint64	`gorm:"column:sequence;not null;default: 5;comment:'排序值'" json:"sequence" form:"sequence"`
	//ParentId	uint64	`gorm:"column:parent_id;not null;comment:'父级ID'" json:"parent_id" form:"parent_id"`
}

func (Role) TableName() string {
	return model.TableName("role")
}

// 添加之前
func (m *Role) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新之前
func (m *Role) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

