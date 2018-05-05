package routers

import (
	"github.com/astaxie/beego"
	"isoft_sso_web/controllers"
)

func init() {
	beego.Router("/user/regist", &controllers.UserController{}, "get,post:Regist")
	beego.Router("/user/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/user/checkOrInValidateTokenString", &controllers.UserController{}, "get,post:CheckOrInValidateTokenString")

	beego.Router("/userlogin/loginRecordList", &controllers.LoginRecordController{}, "get,post:LoginRecordList")

	beego.Router("/appregister/appRegisterList", &controllers.AppRegisterController{}, "get,post:AppRegisterList")
	beego.Router("/appregister/addAppRegister", &controllers.AppRegisterController{}, "get,post:AddAppRegister")

}
