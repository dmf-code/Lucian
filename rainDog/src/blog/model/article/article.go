package article

type PostField struct {
	Title string `json:"title"`
	CheckedCategorys string `json:"checkedCategorys" gorm:"column:category_ids;"`
	CheckedTags string `json:"checkedTags" gorm:"column:tag_ids;"`
	MdCode string `json:"mdCode" gorm:"column:md_code;"`
	HtmlCode string `json:"htmlCode" gorm:"html_code;"`
}

type PutField struct {
	Id string `json:"id"`
	Title string `json:"title"`
	CheckedCategorys string `json:"checked_categorys"`
	CheckedTags string `json:"checked_tags"`
	MdCode string `json:"md_code"`
	HtmlCode string `json:"html_code"`
}

type GetField struct {
	Id string `json:"id"`
	Title string `json:"title"`
	CheckedCategorys string `json:"checked_categorys" gorm:"column:category_ids;"`
	CheckedTags string `json:"checked_tags" gorm:"column:tag_ids;"`
	MdCode string `json:"mdCode"`
	HtmlCode string `json:"htmlCode"`
}

type DeleteField struct {
	Id string `json:"id"`
}
