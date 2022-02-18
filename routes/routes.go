/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 20:14:08
 * @LastEditTime: 2022-01-23 18:09:18
 */
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mygin/controllers"
)

func InitRoutes(app *gin.Engine) {

	//测试
	app.GET("/testing/:id", controllers.StartPage)
	app.POST("/testing", controllers.StartPage)
	//登录
	app.POST("/login", controllers.Login)
	//管理员
	app.GET("/auth", controllers.AdminAuth)
	app.GET("/admin", controllers.AdminList)
	app.GET("/admin/:id", controllers.AdminGet)
	app.POST("/admin", controllers.AdminCreate)
	app.PUT("/admin/:id", controllers.AdminUpdate)
	app.DELETE("/admin/:id", controllers.AdminDelete)

	//重定义NoMethod  NoRoute
	app.NoMethod(controllers.HandleNoMethod)
	app.NoRoute(controllers.HandleNoRoute)
}
