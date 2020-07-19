package Table

import (
	"app/model"
	"app/library/format"
	"github.com/jinzhu/gorm"
	"time"
)

type Admin struct {
	model.BaseModel
	Username string `gorm:"column:username;size:32;not null;unique_index; comment:'用户名';" json:"username" form: "username"`
	Password string `gorm:"column:password;size:128;not null;" json:"password; comment:'密码';" form: "password"`
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
