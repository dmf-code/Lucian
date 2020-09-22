package routes

import (
	"github.com/gin-gonic/gin"
	"rain/internal/app/model"
)

func Backend(r *gin.RouterGroup) {
	categoryGroup(r)

	tagGroup(r)

	articleGroup(r)

	adminGroup(r)

	roleGroup(r)

	menuGroup(r)

	adminRoleGroup(r)

	roleMenuGroup(r)

	tutorialGroup(r)
}

func categoryGroup(r *gin.RouterGroup) {
	category := model.Category{}
	r.POST("/category", category.Store)

	r.PUT("/category/:id", category.Update)

	r.GET("/category", category.Index)

	r.DELETE("/category/:id", category.Destroy)
}

func tagGroup(r *gin.RouterGroup) {
	tag := model.Tag{}
	r.POST("/tag", tag.Store)

	r.PUT("/tag/:id", tag.Update)

	r.GET("/tag", tag.Index)

	r.DELETE("/tag/:id", tag.Destroy)
}

func articleGroup(r *gin.RouterGroup) {
	article := model.Article{}
	r.POST("/article", article.Store)

	r.PUT("/article/:id", article.Update)

	r.GET("/article", article.Index)

	r.GET("/article/:id", article.Show)

	r.DELETE("/article/:id", article.Destroy)
}

func adminGroup(r *gin.RouterGroup) {
	admin := model.Admin{}
	r.POST("/admin", admin.Store)

	r.PUT("/admin/:id", admin.Update)

	r.GET("/admin", admin.Index)

	r.GET("/admin/:id", admin.Show)

	r.DELETE("/admin/:id", admin.Destroy)

	r.PUT("/resetPassword/:id", admin.ResetPassword)
}

func roleGroup(r *gin.RouterGroup) {
	role := model.Role{}
	r.POST("/role", role.Store)

	r.PUT("/role/:id", role.Update)

	r.GET("/role", role.Index)

	r.GET("/role/:id", role.Show)

	r.DELETE("/role/:id", role.Destroy)
}

func menuGroup(r *gin.RouterGroup) {
	menu := model.Menu{}
	r.POST("/menu", menu.Store)

	r.PUT("/menu/:id", menu.Update)

	r.GET("/menu", menu.Index)

	r.GET("/menu/:id", menu.Show)

	r.DELETE("/menu/:id", menu.Destroy)

	r.GET("/menuList", menu.List)

	r.GET("/menuApiList", menu.ApiList)
}

func adminRoleGroup(r *gin.RouterGroup) {
	adminRole := model.AdminRole{}
	r.POST("/adminRole", adminRole.Store)

	r.PUT("/adminRole/:id", adminRole.Update)

	r.GET("/adminRole", adminRole.Index)

	r.GET("/adminRole/:id", adminRole.Show)

	r.DELETE("/adminRole/:id", adminRole.Destroy)
}

func roleMenuGroup(r *gin.RouterGroup) {
	roleMenu := model.RoleMenu{}
	r.POST("/roleMenu", roleMenu.Store)

	r.PUT("/roleMenu/:id", roleMenu.Update)

	r.GET("/roleMenu", roleMenu.Index)

	r.GET("/roleMenu/:id", roleMenu.Show)

	r.DELETE("/roleMenu/:id", roleMenu.Destroy)

	r.GET("/roleMenuList", roleMenu.List)
}

func tutorialGroup(r *gin.RouterGroup) {
	tutorial := model.Tutorial{}
	contentTutorial := model.ContentTutorial{}
	r.POST("/tutorial", tutorial.Store)

	r.PUT("/tutorial/:id", tutorial.Update)

	r.GET("/tutorial", tutorial.Index)

	r.GET("/tutorial/:id", tutorial.Show)

	r.DELETE("/tutorial/:id", tutorial.Destroy)

	r.GET("/tutorialList/:pid", contentTutorial.List)
}
