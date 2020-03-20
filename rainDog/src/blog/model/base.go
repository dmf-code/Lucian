package model

import (
	"blog/utils/format"
	"fmt"
)

type BaseModel struct {
	ID        uint             `gorm:"primary_key" json:"id"`
	CreatedAt format.JSONTime  `json:"createdAt"`
	UpdatedAt format.JSONTime  `json:"updatedAt"`
	DeletedAt *format.JSONTime `sql:"index" json:"-"`
}

const tablePrefix = ""

func TableName(name string) string {
	return fmt.Sprintf("%s%s%s", tablePrefix, "", name)
}
