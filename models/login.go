/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 21:38:19
 * @LastEditTime: 2022-01-23 18:33:50
 */
package models

//https://raw.githubusercontent.com/go-playground/validator/v9/_examples/translations/main.go  规则例子
type Login struct {
	Username string `json:"username" form:"username" binding:"required,lt=20,gt=6" zh:"用户名"`
	Password string `json:"password" form:"password" binding:"required,lt=20,gt=5" zh:"密码"`
	//不能为空并且大于10
	// Age      int       `form:"age" binding:"required,gt=10,gte=0,lte=130"`
	// Email    string    `form:"email" binding:"required,email"`
	// Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}
