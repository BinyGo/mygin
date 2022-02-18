/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 18:31:40
 * @LastEditTime: 2022-02-10 22:46:13
 */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mygin/middleware"
	"github.com/mygin/routes"
	"github.com/mygin/validator"
)

func main() {
	validator.Init()
	//models.Init()

	app := gin.Default()

	middleware.InitMiddleware(app)
	routes.InitRoutes(app)

	app.Run("127.0.0.1:8085")

}
