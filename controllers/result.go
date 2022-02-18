/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-23 10:16:19
 * @LastEditTime: 2022-01-24 19:31:26
 */
package controllers

import "github.com/gin-gonic/gin"

// type ResultJson struct {
// 	Code  int         `json:"code"`
// 	Msg   string      `json:"msg"`
// 	Data  interface{} `json:"data"`
// 	Token string      `json:"token"`
// }

const (
	failedCode  = 0
	successCode = 1
)

type resultJson struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`  //omitempty返回信息时是零值，就会忽略这个字段
	Token string      `json:"token,omitempty"` //omitempty返回信息时是零值，就会忽略这个字段
}

type resultError struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

func HandleNoMethod(c *gin.Context) {
	result := &resultError{
		Code: failedCode,
		Msg:  "Method Not Allowed",
	}
	c.JSON(200, result)
	//return
}

func HandleNoRoute(c *gin.Context) {
	result := &resultError{
		Code: failedCode,
		Msg:  "Route Not Found",
	}
	c.JSON(200, result)
	//return
}
