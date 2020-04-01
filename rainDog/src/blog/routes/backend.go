package routes

import (
	"blog/model/article"
	"blog/model/category"
	"blog/model/menu"
	"blog/model/role"
	"blog/model/roleMenu"
	"blog/model/tag"
	"github.com/gin-gonic/gin"
)

func Backend(r *gin.RouterGroup) {
	categoryGroup(r)

	tagGroup(r)
	
	articleGroup(r)

	roleGroup(r)

	menuGroup(r)

	roleMenuGroup(r)
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

func menuGroup(r *gin.RouterGroup) {
	r.POST("/menu", menu.Store)

	r.PUT("/menu/:id", menu.Update)

	r.GET("/menu", menu.Index)

	r.GET("/menu/:id", menu.Show)

	r.DELETE("/menu/:id", menu.Destroy)
}

func roleMenuGroup(r *gin.RouterGroup) {
	r.POST("/roleMenu", roleMenu.Store)

	r.PUT("/roleMenu/:id", roleMenu.Update)

	r.GET("/roleMenu", roleMenu.Index)

	r.GET("/roleMenu/:id", roleMenu.Show)

	r.DELETE("/roleMenu/:id", roleMenu.Destroy)
}


