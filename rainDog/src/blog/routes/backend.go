package routes

import (
	"blog/model/category"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Backend(r *gin.RouterGroup) {
	r.POST("/category", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var field category.Field
		_ = context.BindJSON(&field)
		db.Table("category").Create(&field)
		helper.Success(context, 200, gin.H{"id": field.Id})
	})

	r.GET("/category", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var fields []category.Field
		db.Table("category").Find(&fields)
		fmt.Println(fields)
		helper.Success(context, 200, gin.H{"list": fields})
	})
}
