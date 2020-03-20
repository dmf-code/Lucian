package manage

import (
	"blog/model"
	"blog/utils/format"
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	model.BaseModel
	Name		string	`gorm:"column:name;size:32;not null;" json:"name" form:"name"`			// 名称
	Memo 		string	`gorm:"column:memo;size:64;" json:"memo" form:"memo"`					// 备注
	Sequence	string	`gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`		// 排序值
	ParentId	uint64	`gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"`	// 父级ID
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
func (m *Role) BeforeUpdate(scope gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}
