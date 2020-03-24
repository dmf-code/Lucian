package manage

import (
	"blog/model"
	"blog/utils/format"
	"github.com/jinzhu/gorm"
	"time"
)

type AdminRole struct {
	model.BaseModel
	AdminId		uint64	`gorm:"column:admin_id;unique_index:uk_admin_role_admin_id;not null;'"`		// 管理员ID
	RoleId		uint64	`gorm:"column:role_id;unique_index:uk_admin_role_admin_id;not null;"`		// 角色ID
}

// 表名
func (AdminRole) TableName() string {
	return model.TableName("admin_role")
}

// 添加前
func (m *AdminRole) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新前
func (m *AdminRole) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

