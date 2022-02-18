/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-23 09:20:15
 * @LastEditTime: 2022-01-24 19:24:12
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mygin/util"
)

type resultJson struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`  //omitempty返回信息时是零值，就会忽略这个字段
	Token string      `json:"token,omitempty"` //omitempty返回信息时是零值，就会忽略这个字段
}

func checkToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		skip := c.FullPath()
		if skip != "/login" {
			result := &resultJson{Code: 1}

			//获取token
			token := c.Request.Header.Get("Authorization")
			//result.Token = token
			if token == "" {
				//错误返回JSON格式信息
				result.Msg = "token不能为空"
				c.JSON(200, result)
				c.Abort()
				return
			}
			token = token[7:] //去除"Bearer "token头
			//验证jwt token
			payload, err := util.ValidateToken(token)
			if err != nil {
				//错误返回JSON格式信息
				result.Code = -1       //-1为登录超时,需要前端判断跳转到登录页面
				result.Msg = "token无效" //err.Error()
				c.JSON(200, result)
				c.Abort()
				return
			}
			c.Set("uid", payload.UserID)
		}
	}
}
