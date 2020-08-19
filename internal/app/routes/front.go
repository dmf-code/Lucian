package routes

import (
	"app/model/article"
	"github.com/gin-gonic/gin"
)

func Front(r *gin.RouterGroup)  {
	ArticleGroup(r)
	tutorialGroup(r)
}

func ArticleGroup(r *gin.RouterGroup)  {
	r.GET("/article", article.Index)

	r.GET("/article/:id", article.Show)

}