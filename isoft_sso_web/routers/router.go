package routers

import (
	"isoft_sso_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/regist", &controllers.UserController{}, "get,post:Regist")
	beego.Router("/user/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/user/logout", &controllers.UserController{}, "get,post:Logout")
	beego.Router("/user/checkLogin", &controllers.UserController{}, "get,post:CheckLogin")
	beego.Router("/user/deleteToken", &controllers.UserController{}, "get,post:DeleteToken")

	beego.Router("/userlogin/loginRecordList", &controllers.LoginRecordController{}, "get,post:LoginRecordList")

	beego.Router("/appregister/appRegisterList", &controllers.AppRegisterController{}, "get,post:AppRegisterList")
	beego.Router("/appregister/addAppRegister", &controllers.AppRegisterController{}, "get,post:AddAppRegister")

}
