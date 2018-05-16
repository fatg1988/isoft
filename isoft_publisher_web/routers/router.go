package routers

import (
	"github.com/astaxie/beego"
	"isoft_publisher_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
