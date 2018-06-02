package routers

import (
	"github.com/astaxie/beego"
	"github.com/isoft/isoft_sso_web/controllers"
)

func init() {
	beego.Router("/user/regist", &controllers.UserController{}, "get,post:Regist")
	beego.Router("/user/login", &controllers.UserController{}, "get,post:Login")

	beego.Router("/userlogin/loginRecordList", &controllers.LoginRecordController{}, "get,post:LoginRecordList")

	beego.Router("/appregister/appRegisterList", &controllers.AppRegisterController{}, "get,post:AppRegisterList")
	beego.Router("/appregister/addAppRegister", &controllers.AppRegisterController{}, "get,post:AddAppRegister")

	// sso 简单认证模型,每次请求都会在登录系统进行认证,客户端不进行任何认证操作
	beego.Router("/user/checkOrInValidateTokenString", &controllers.UserController{}, "get,post:CheckOrInValidateTokenString")
}
