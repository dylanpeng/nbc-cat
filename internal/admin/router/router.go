package router

import (
	ctrl "cat/common/control"
	"cat/common/middleware"
	"cat/internal/admin/login/control"
	"github.com/gin-gonic/gin"
)

var Router = &router{}

type router struct{}

func (r *router) GetIdentifier(ctx *gin.Context) string {
	// do author return id or something can identify
	return "unknown"
}

func (r *router) RegHttpHandler(app *gin.Engine) {
	app.Any("/health", ctrl.Health)
	app.Use(middleware.CheckEncoding)
	app.Use(middleware.CrossDomain)

	manageGroup := app.Group("/manage")
	{
		manageGroup.POST("/login", control.AdminUser.Login)
	}

	manageUserGroup := app.Group("/manage/user")
	{
		//manageUserGroup.GET("/list", control.AdminPerm.Tree)
		manageUserGroup.POST("/add", control.AdminUser.Add)
		//manageUserGroup.GET("/edit", control.AdminUser.Logout)
		//manageUserGroup.GET("/edit", control.AdminUser.Logout)
	}

}
