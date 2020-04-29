package manage

import (
	"blog/model"
	"blog/utils/format"
	"github.com/jinzhu/gorm"
	"time"
)

type Admin struct {
	model.BaseModel
	Username string `gorm:"column:username;size:32;not null;unique_index;" json:"username" form: "username"`			// 用户名
	Password string `gorm:"column:password;size:32;not null;" json:"password" form: "password"`							// 密码
}

func (Admin) TableName()string {
	return model.TableName("admin")
}

// 添加之前
func (m *Admin) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = format.JSONTime{Time: time.Now()}
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}

// 更新之前
func (m *Admin) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = format.JSONTime{Time: time.Now()}
	return nil
}
