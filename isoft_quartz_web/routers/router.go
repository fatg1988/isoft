package routers

import (
	"github.com/astaxie/beego"
	"isoft_quartz_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
