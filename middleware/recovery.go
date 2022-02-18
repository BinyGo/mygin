/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-24 19:21:45
 * @LastEditTime: 2022-01-24 19:47:16
 */
package middleware

import (
	"github.com/gin-gonic/gin"
)

// recovery 机制，将协程中的函数异常进行捕获

func recovery() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		// 核心在增加这个 recover 机制，捕获 c.Next()出现的 panic
		defer func() {
			if err := recover(); err != nil {
				result := make(map[string]interface{})
				result["code"] = 0
				result["msg"] = "服务器异常"
				result["error"] = err
				c.JSON(200, result)
			}
		}()
		// 使用 next 执行具体的业务逻辑
		c.Next()
	}
}
