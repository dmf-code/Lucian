package tutorial

import "app/model"

type CoverTutorial struct {
	model.BaseModel
	Img  		string		`json:"img" gorm:"size:256; column:img; comment:'教程封面图片'"`
	Title		string		`json:"title" gorm:"size:64; column:title; comment: '标题';"`
	Root		int			`json:"root" gorm:"column:root; comment: '目录根节点';"`
}
