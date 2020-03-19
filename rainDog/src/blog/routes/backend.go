package routes

import (
	"blog/model/article"
	"blog/model/category"
	"blog/model/tag"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Backend(r *gin.RouterGroup) {
	categoryGroup(r)

	tagGroup(r)
	
	articleGroup(r)
}

func categoryGroup(r *gin.RouterGroup) {
	r.POST("/category", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var field category.PostField
		_ = context.BindJSON(&field)
		if err := db.Table("category").Create(&field).Error; err != nil {
			helper.Fail(context, err.Error())
			return
		}
		helper.Success(context,  "success")
	})

	r.PUT("/category/:id/:name", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var putField category.PutField
		putField.Id = context.Param("id")
		putField.Name = context.Param("name")
		fmt.Println(putField)
		if err := db.Table("category").Model(&putField).Update("name", putField.Name).Error; err != nil {
			helper.Fail(context, err.Error())
			return
		}

		helper.Success(context, "success")
	})

	r.GET("/category", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var fields []category.GetField
		if err := db.Table("category").Find(&fields).Error; err != nil {
			helper.Fail(context, err.Error())
			return
		}
		fmt.Println(fields)
		helper.Success(context, fields)
	})

	r.DELETE("/category/:id", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var field category.DeleteField
		field.Id = context.Param("id")
		if err := db.Table("category").Delete(&field).Error; err != nil {
			helper.Fail(context, err.Error())
			return
		}
		helper.Success(context, "success")
	})
}

func tagGroup(r *gin.RouterGroup) {
	r.POST("/tag", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var field tag.PostField
		_ = context.BindJSON(&field)
		if err := db.Table("tag").Create(&field).Error; err != nil {
			helper.Fail(context, err.Error())
			return
		}
		helper.Success(context, "success")
	})

	r.PUT("/tag/:id/:name", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var putField tag.PutField
		putField.Id = context.Param("id")
		putField.Name = context.Param("name")
		fmt.Println(putField)
		if err := db.Table("tag").Model(&putField).Update("name", putField.Name).Error; err != nil {
			helper.Fail(context, err.Error())
			return
		}

		helper.Success(context,"success")
	})

	r.GET("/tag", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var fields []tag.GetField
		if err := db.Table("tag").Find(&fields).Error; err != nil {
			helper.Fail(context, err.Error())
			return
		}
		fmt.Println(fields)
		helper.Success(context, fields)
	})

	r.DELETE("/tag/:id", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var field tag.DeleteField
		field.Id = context.Param("id")
		if err := db.Table("tag").Delete(&field).Error; err != nil {
			helper.Fail(context, err.Error())
			return
		}
		helper.Success(context, "success")
	})
}

func articleGroup(r *gin.RouterGroup) {
	r.POST("/article", article.Store)
	
	r.PUT("/article/:id", article.Update)
	
	r.GET("/article", article.Index)

	r.GET("/article/:id", article.Show)

	r.DELETE("/article/:id", article.Destroy)
}
