package routers

import (
	"github.com/astaxie/beego"
	"isoft_timer_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
