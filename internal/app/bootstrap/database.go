package bootstrap

import (
	"app/bootstrap/Table"
	"app/library/helper"
	"app/model/article"
	"app/model/category"
	"app/model/tag"
	"app/model/tutorial"
	"fmt"
)


func InitTable() {
	db := helper.Db()
	fmt.Println(db.AutoMigrate(new(Table.Menu)).Error)
	fmt.Println(db.AutoMigrate(new(Table.Role)).Error)
	fmt.Println(db.AutoMigrate(new(Table.Admin)).Error)
	fmt.Println(db.AutoMigrate(new(Table.RoleMenu)).Error)
	fmt.Println(db.AutoMigrate(new(Table.AdminRole)).Error)
	fmt.Println(db.AutoMigrate(new(article.Article)).Error)
	fmt.Println(db.AutoMigrate(new(tag.Tag)).Error)
	fmt.Println(db.AutoMigrate(new(category.Category)).Error)
	fmt.Println(db.AutoMigrate(new(tutorial.CoverTutorial)).Error)
}