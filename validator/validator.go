/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 22:02:23
 * @LastEditTime: 2022-01-24 17:08:58
 */
package validator

//gin > 1.4.0

//将验证器错误翻译成中文

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni *ut.UniversalTranslator
	//validate *validator.Validate
	trans ut.Translator
)

func Init() {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	//增加手机号码验证规则
	validate.RegisterValidation("mobile", checkMobile)
	//增加手机号码验证中文信息
	validate.RegisterTranslation("mobile", trans, registerTranslator("mobile", "{0}格式错误"), diyTranslate)
	//结构体字段zh中文名 解决field无法翻译,出现 "Email必须是一个有效的邮箱" 的问题,结构体中增加zh中文值,进行替换
	validate.RegisterTagNameFunc(validateZhName)
	//自定义结构体验证规则? 确认密码?
	//validate.RegisterStructValidation(rePassword, models.Admin{})
	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)

}

//结构体字段zh中文名 解决field无法翻译,出现 "Email必须是一个有效的邮箱" 的问题,结构体中增加zh中文值,进行替换
func validateZhName(fld reflect.StructField) string {
	//fmt.Println("-----------", fld.Tag.Get("zh"))
	name := strings.SplitN(fld.Tag.Get("zh"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// // diyTranslate 自定义字段的翻译方法
func diyTranslate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

//Translate 翻译错误信息
func Translate(err error) string {

	var result = ""
	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		if result == "" {
			result += err.Translate(trans)
		} else {
			result += "," + err.Translate(trans)
		}
	}
	return result
}

//增加手机号码验证规则
var checkMobile validator.Func = func(fl validator.FieldLevel) bool {
	//fmt.Println("==========", fl.Parent().FieldByName("Password"), "=========")
	if mobile, ok := fl.Field().Interface().(string); ok {
		//regular := `^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$`
		regular := `^(1[3-9])\d{9}$`
		reg := regexp.MustCompile(regular)
		return reg.MatchString(mobile)
	}
	return true
}

// //备份Translate代码
// func Translate2(err error) map[string][]string {

// 	var result = make(map[string][]string)

// 	errors := err.(validator.ValidationErrors)

// 	for _, err := range errors {
// 		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
// 	}
// 	return result
// }

// 增加确认密码验证规则
// func rePassword(sl validator.StructLevel) {
// 	fmt.Println("==========", sl, "=========")
// 	su := sl.Current().Interface().(models.Admin)
// 	if su.Password != su.RePassword {
// 		// 输出错误提示信息，最后一个参数就是传递的param
// 		sl.ReportError(su.RePassword, "确认密码", "rePassword", "eqfield", "密码")
// 	}
// }
