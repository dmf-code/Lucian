package manage

import (
	"app/model"
	"app/utils/format"
	"github.com/jinzhu/gorm"
	"time"
)

type RoleMenu struct {
	model.BaseModel
	RoleId	uint64	`gorm:"column:role_id;unique_index:uk_role_menu_role_id;not null;"`		// 角色ID
	MenuId	uint64	`gorm:"column:menu_id;unique_index:uk_role_menu_role_id;not null;"`		// 菜单ID
}

func (RoleMenu) TableName() string {
	return model.TableName("role_menu")
}
// 添加前
func (m *RoleMenu) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新前
func (m *RoleMenu) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}