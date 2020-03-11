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
		var field category.PostField
		_ = context.BindJSON(&field)
		if err := db.Table("category").Create(&field).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}
		helper.Success(context, 200, gin.H{"msg": "success"})
	})

	r.PUT("/category/:id/:name", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var putField category.PutField
		putField.Id = context.Param("id")
		putField.Name = context.Param("name")
		fmt.Println(putField)
		if err := db.Table("category").Model(&putField).Update("name", putField.Name).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}

		helper.Success(context, 200, gin.H{"msg": "success"})
	})

	r.GET("/category", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var fields []category.GetField
		if err := db.Table("category").Find(&fields).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}
		fmt.Println(fields)
		helper.Success(context, 200, gin.H{"list": fields})
	})
}
