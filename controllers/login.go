/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 20:29:59
 * @LastEditTime: 2022-01-23 19:18:33
 */
package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mygin/models"
	"github.com/mygin/util"
	"github.com/mygin/validator"
)

func Login(c *gin.Context) {
	result := &resultJson{Code: failedCode, Msg: "登录失败"}
	login := &models.Login{}

	if err := c.ShouldBind(login); err != nil {
		result.Msg = validator.Translate(err)
		c.JSON(200, result)
		return
	}

	//数据库查询数据
	admin := models.GetAdminByUsername(login.Username)
	if admin.Id == 0 {
		result.Msg = "用户不存在"
	} else if admin.Password != util.Password(login.Password) {
		result.Msg = "密码错误"
	} else {
		token, err := util.GenerateToken(admin.Id, admin.Username, 60*24*24)
		if err != nil {
			result.Msg = err.Error()
		} else {
			result.Code = successCode
			result.Msg = "登录成功"
			//result.Data = admin
			result.Token = token

		}
	}
	c.JSON(200, result)
}

func StartPage(c *gin.Context) {
	//var person Person
	person := &models.Person{} //新建
	person1 := person          //引用赋值
	person2 := *person         //复制赋值
	person1.Id = 11
	person2.Name = "22"
	fmt.Println(person)  //&{11 }
	fmt.Println(person1) //&{11 }
	fmt.Println(person2) //{0 22}

	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
	id := c.Param("id")
	if c.ShouldBind(person) == nil {
		fmt.Println(person)
		log.Println(id)
		log.Println(person.Id)
		log.Println(person.Name)
	}

	c.String(200, "Success")
}
