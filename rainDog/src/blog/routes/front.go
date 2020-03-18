package routes

import (
	"blog/model/article"
	"blog/utils/helper"
	"github.com/gin-gonic/gin"
)

func Front(r *gin.RouterGroup)  {
	ArticleGroup(r)
}

func ArticleGroup(r *gin.RouterGroup)  {
	r.GET("/article", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var fields []article.GetField
		if err := db.Table("article").Find(&fields).Error; err != nil {
			helper.Fail(context, "查询失败")
			return
		}

		helper.Success(context, fields)
	})

	r.GET("/article/:id", func(context *gin.Context) {
		db := helper.Db("rain_dog")
		var field article.GetField
		if err := db.Table("article").Where("id = ?", context.Param("id")).First(&field).Error; err != nil {
			helper.Fail(context, "查询失败")
			return
		}

		helper.Success(context, field)
	})

}