/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-23 11:42:27
 * @LastEditTime: 2022-02-10 22:38:12
 */
package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(app *gin.Engine) {
	app.Use(cors())
	app.Use(checkToken())
	app.Use(recovery())
}
