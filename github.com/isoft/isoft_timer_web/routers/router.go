package routers

import (
	"github.com/astaxie/beego"
	"github.com/isoft/isoft_timer_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
