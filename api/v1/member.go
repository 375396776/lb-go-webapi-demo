/*
 description:

 @author lib
 @since 2020/05/13
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"invoice/common/response"
	"invoice/log"
	member2 "invoice/model/member"
	"regexp"
	"strings"
)

func AddMember(ctx *gin.Context) {
	var member member2.Member
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("phoneValid", phoneValid)
	}
	if err := ctx.ShouldBindWith(&member, binding.JSON); err != nil {
		log.Error.Println(err)
		if strings.Contains(err.Error(), "Member.Mobile") {
			response.FailResult(400, "手机号格式错误", ctx)
			return
		}
		if strings.Contains(err.Error(), "Member.IdCard") {
			response.FailResult(400, "身份证号格式错误", ctx)
			return
		}
		if strings.Contains(err.Error(), "Member.Name") {
			response.FailResult(400, "姓名不能为空", ctx)
			return
		}
		if strings.Contains(err.Error(), "Member.Age") {
			response.FailResult(400, "年龄要在10，120之间", ctx)
			return
		}
	}
	response.SuccessResult(&member, ctx)
}

var phoneValid validator.Func = func(fl validator.FieldLevel) bool {
	//手机号正则
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	s := fl.Field().Interface().(string)
	compile := regexp.MustCompile(regular)
	if compile.MatchString(s) {
		return true
	}
	return false
}
