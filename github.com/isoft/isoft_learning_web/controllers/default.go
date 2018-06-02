package controllers

import (
	"github.com/astaxie/beego"
	"github.com/isoft/isoft_learning_web/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

func (this *MainController) Index() {
	coursetypelist := this.GetString("coursetypelist")
	if coursetypelist == "coursetypelist" {
		this.Data["CourseTypeListShow"] = "CourseTypeListShow"
	}

	// 热门推荐,根据观看量查询前 50 个
	condArr := make(map[string]string)
	condArr["querysOrder"] = "-watch_number"
	courses, _, err := models.QueryCourse(condArr, 1, 50)
	if err == nil {
		this.Data["Recommends"] = courses
	}

	this.TplName = "index.html"
}
