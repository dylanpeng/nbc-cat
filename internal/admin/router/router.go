package router

import (
	ctrl "cat/common/control"
	commonMiddleware "cat/common/middleware"
	"cat/internal/admin/logic/control"
	"cat/internal/admin/logic/middleware"
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
	app.Use(commonMiddleware.CheckEncoding)
	app.Use(commonMiddleware.CrossDomain)

	manageGroup := app.Group("/manage")
	{
		manageGroup.POST("/login", control.AdminUser.Login)
	}

	userGroup := app.Group("/manage/user", middleware.AdminAuth, middleware.PermissionCheck)
	{
		//userGroup.GET("/list", control.AdminPerm.Tree)
		userGroup.POST("/add", control.AdminUser.Add)
		//userGroup.GET("/edit", control.AdminUser.Logout)
		//userGroup.GET("/edit", control.AdminUser.Logout)
	}

}
