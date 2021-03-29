package handler

import (
	"github.com/gin-gonic/gin"
	"rain/internal/model"
)

func FrontHandler(r *gin.RouterGroup) {
	frontArticleGroup(r)
	frontTutorialGroup(r)

	frontCategoryGroup(r)

	navGroup(r)

	bookmarkGroup(r)
}

func frontArticleGroup(r *gin.RouterGroup) {
	article := model.Article{}
	r.GET("/article", article.Index)

	r.GET("/article/:id", article.Show)

}

func frontTutorialGroup(r *gin.RouterGroup) {
	tutorial := model.Tutorial{}
	contentTutorial := model.ContentTutorial{}
	r.GET("/tutorial", tutorial.Index)

	r.GET("/tutorial/:id", tutorial.Show)

	r.GET("/tutorialContent/:id", contentTutorial.Show)

	r.GET("/tutorialList/:pid", contentTutorial.List)
}

func frontCategoryGroup(r *gin.RouterGroup) {
	category := model.Category{}
	r.GET("/category", category.Index)
}
