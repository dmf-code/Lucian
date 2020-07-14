package tutorial

import "app/model"

type ContentTutorial struct {
	model.BaseModel
	MdCode		string `json:"mdCode" gorm:"type:longtext;column:md_code;comment:'markdown代码'"`
	HtmlCode	string `json:"htmlCode" gorm:"type:longtext;column:html_code;comment:'html代码'"`
}
