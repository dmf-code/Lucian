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

	r.DELETE("/category/:id", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var field category.DeleteField
		field.Id = context.Param("id")
		if err := db.Table("category").Delete(&field).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}
		helper.Success(context, 200, gin.H{"msg": "success"})
	})
}

func tagGroup(r *gin.RouterGroup) {
	r.POST("/tag", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var field tag.PostField
		_ = context.BindJSON(&field)
		if err := db.Table("tag").Create(&field).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}
		helper.Success(context, 200, gin.H{"msg": "success"})
	})

	r.PUT("/tag/:id/:name", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var putField tag.PutField
		putField.Id = context.Param("id")
		putField.Name = context.Param("name")
		fmt.Println(putField)
		if err := db.Table("tag").Model(&putField).Update("name", putField.Name).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}

		helper.Success(context, 200, gin.H{"msg": "success"})
	})

	r.GET("/tag", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var fields []tag.GetField
		if err := db.Table("tag").Find(&fields).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}
		fmt.Println(fields)
		helper.Success(context, 200, gin.H{"list": fields})
	})

	r.DELETE("/tag/:id", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var field tag.DeleteField
		field.Id = context.Param("id")
		if err := db.Table("tag").Delete(&field).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}
		helper.Success(context, 200, gin.H{"msg": "success"})
	})
}

func articleGroup(r *gin.RouterGroup) {
	r.POST("/article", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var field article.PostField
		fmt.Println(context.PostForm("mdCode"))
		context.BindJSON(&field)
		fmt.Println(field)

		if err := db.Table("article").Create(&field).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}

		helper.Success(context, 200, gin.H{"msg": "success"})
	})
	
	r.PUT("/article/:id", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var filed article.PutField
		context.BindJSON(&filed)
		filed.Id = context.Param("id")
		fmt.Println(filed)
		if err := db.Table("article").Model(&filed).Updates(filed).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}

		helper.Success(context, 200, gin.H{"msg": "更新成功"})
	})
	
	r.GET("/article", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var fields []article.GetField
		if err := db.Table("article").Find(&fields).Error; err != nil {
			helper.Fail(context, 200, "查询失败")
			return
		}

		helper.Success(context, 200, gin.H{"list": fields})
	})

	r.GET("/article/:id", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var field article.GetField
		if err := db.Table("article").First(&field).Error; err != nil {
			helper.Fail(context, 200, "查询失败")
			return
		}

		helper.Success(context, 200, gin.H{"list": field})
	})

	r.DELETE("/article/:id", func(context *gin.Context) {
		db, _ := helper.Db("rain_dog")
		var field article.DeleteField
		field.Id = context.Param("id")
		if err := db.Table("article").Delete(&field).Error; err != nil {
			helper.Fail(context, 200, err.Error())
			return
		}

		helper.Success(context, 200, gin.H{"msg": "删除成功"})
	})
}
