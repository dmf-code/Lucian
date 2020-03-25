package routes

import (
	"blog/model/article"
	"blog/model/category"
	"blog/model/role"
	"blog/model/tag"
	"github.com/gin-gonic/gin"
)

func Backend(r *gin.RouterGroup) {
	categoryGroup(r)

	tagGroup(r)
	
	articleGroup(r)

	roleGroup(r)
}

func categoryGroup(r *gin.RouterGroup) {
	r.POST("/category", category.Store)

	r.PUT("/category/:id", category.Update)

	r.GET("/category", category.Index)

	r.DELETE("/category/:id", category.Destroy)
}

func tagGroup(r *gin.RouterGroup) {
	r.POST("/tag", tag.Store)

	r.PUT("/tag/:id", tag.Update)

	r.GET("/tag", tag.Index)

	r.DELETE("/tag/:id", tag.Destroy)
}

func articleGroup(r *gin.RouterGroup) {
	r.POST("/article", article.Store)
	
	r.PUT("/article/:id", article.Update)
	
	r.GET("/article", article.Index)

	r.GET("/article/:id", article.Show)

	r.DELETE("/article/:id", article.Destroy)
}

func roleGroup(r *gin.RouterGroup) {
	r.POST("/role", role.Store)

	r.PUT("/role/:id", role.Update)

	r.GET("/role", role.Index)

	r.GET("/role/:id", role.Show)

	r.DELETE("/role/:id", role.Destroy)
}
