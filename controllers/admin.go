/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 20:37:41
 * @LastEditTime: 2022-01-24 22:23:00
 */
package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mygin/models"
	"github.com/mygin/util"
	"github.com/mygin/validator"
)

type ResultJson struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Token string      `json:"token,omitempty"`
}

func AdminGet(c *gin.Context) {
	//fmt.Println(c.Params)
	result := &ResultJson{Code: failedCode, Msg: "用户不存在"}

	id := c.Param("id")
	nid, err := strconv.ParseInt(id, 10, 64)
	if err != nil || nid <= 0 {
		c.JSON(200, result)
		return
	}

	admin := models.GetAdminById(nid)
	if admin.Id > 0 {
		result.Msg = "success"
		result.Data = admin
		result.Code = successCode
	}

	c.JSON(200, result)
}

// @Tags 管理员
// @Summary 分页获取管理员
// @Produce  json
// @Param ReqAdmin body models.Admin false "admin"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin [get]
func AdminList(c *gin.Context) {
	result := &ResultJson{Code: failedCode, Msg: "没有更多数据了..."}
	type param struct {
		Limit int `form:"limit" binding:"required,gte=9,lt=2000" default:"10"`
		Page  int `form:"page" binding:"required,gt=0,lt=9000000" default:"1"`
	}
	list := &param{}
	list.Limit = c.GetInt("limit")
	list.Page = c.GetInt("page")
	//c.ShouldBind接收后会清除接收数据,无法二次接收,使用c.ShouldBindBodyWith(&param,binding.JSON)可多次接收
	if err := c.ShouldBind(list); err != nil {
		result.Msg = validator.Translate(err)
		c.JSON(200, result)
		return
	}
	offset := (list.Page - 1) * list.Limit

	admins := models.GetAdminList(list.Limit, offset)
	count := models.GetAdminCount()
	
	if len(admins) > 0 {
		resultdata := make(map[string]interface{})
		resultdata["data"] = admins
		resultdata["total"] = count
		resultdata["current"] = list.Page
		resultdata["limit"] = list.Limit
		result.Code = successCode
		result.Data = resultdata
		result.Msg = "success"
	}

	c.JSON(200, result)
}

func AdminCreate(c *gin.Context) {
	//新增不需要Id,为了跳过验证先设置Id为1
	admin := &models.Admin{Id: 1}
	result := &ResultJson{Code: failedCode, Msg: "添加失败"}

	if err := c.ShouldBind(admin); err != nil {
		result.Msg = validator.Translate(err)
		c.JSON(200, result)
		return
	}
	admin.Id = 0 //验证通过,重新设置Id为0
	//admin.CreateTime = time.Now()
	//admin.UpdateTime = time.Now()
	admin.Password = util.Password(admin.Password)

	err := models.AdminCreate(admin)
	if err != nil {
		result.Msg = err.Error()
	} else {
		result.Code = successCode
		result.Msg = "添加成功"
	}

	c.JSON(200, result)
}

func AdminUpdate(c *gin.Context) {
	//获取数据
	id := c.Param("id")
	admin := &models.Admin{}
	admin.Id, _ = strconv.ParseInt(id, 10, 64)
	result := &ResultJson{Code: failedCode, Msg: "编辑失败"}

	//验证数据
	if err := c.ShouldBind(admin); err != nil {
		result.Msg = validator.Translate(err)
		c.JSON(200, result)
		return
	}
	//调整数据
	//admin.CreateTime = time.Now()
	//admin.UpdateTime = MyTime
	admin.Password = util.Password(admin.Password)
	//更新数据

	ok := models.AdminUpdate(admin)
	if ok {
		result.Code = successCode
		result.Msg = "编辑成功"
	}

	c.JSON(200, result)
}

func AdminDelete(c *gin.Context) {
	result := &ResultJson{Code: failedCode, Msg: "删除失败"}

	id := c.Param("id")
	nid, err := strconv.ParseInt(id, 10, 64)
	if err != nil || nid <= 0 {
		c.JSON(200, result)
		return
	}

	err = models.AdminDelete(nid)
	if err != nil {
		result.Msg = err.Error()
	} else {
		result.Code = successCode
		result.Msg = "删除成功"
	}

	c.JSON(200, result)
}

func AdminAuth(c *gin.Context) {
	result := &ResultJson{Code: failedCode, Msg: "获取失败"}
	uid := c.GetInt64("uid")
	if uid <= 0 {
		c.JSON(200, result)
		return
	}
	resultdata := make(map[string]interface{})
	admin := models.GetAdminById(uid)
	role := models.GetAdminRoles(uid)
	menu := models.GetAllMenu()

	resultdata["admin"] = &admin
	resultdata["role"] = &role
	resultdata["menu"] = &menu

	result.Data = resultdata
	result.Code = successCode
	result.Msg = "获取成功"

	c.JSON(200, result)
}
