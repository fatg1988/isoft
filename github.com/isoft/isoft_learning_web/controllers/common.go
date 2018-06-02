package controllers

import (
	"github.com/astaxie/beego"
	"github.com/isoft/isoft/sso"
	"github.com/isoft/isoft_learning_web/models"
)

type CommonController struct {
	beego.Controller
}

func (this *CommonController) QueryConfiguration() {
	// 获取课程 id
	configuration_name := this.GetString("configuration_name")
	configuration, err := models.QueryConfiguration(configuration_name)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "configuration": configuration}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *CommonController) ToggleFavorite() {
	// 获取课程 id
	favorite_id, _ := this.GetInt("favorite_id")
	favorite_type := this.GetString("favorite_type")
	user_name := this.Ctx.Input.Session("UserName").(string)
	flag := models.IsFavorite(user_name, favorite_id, favorite_type)
	if flag {
		models.DelFavorite(user_name, favorite_id, favorite_type)
	} else {
		models.AddFavorite(user_name, favorite_id, favorite_type)
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func (this *CommonController) CheckLoginUser() {
	if this.GetSession("UserName") != nil {
		this.Data["json"] = &map[string]interface{}{"isLogin": true, "username": this.GetSession("UserName").(string)}
	} else {
		this.Data["json"] = &map[string]interface{}{"isLogin": false}
	}
	this.ServeJSON()
}

func (this *CommonController) Logout() {
	redirectUrl := this.GetString("redirectUrl")
	sso.RedirectToLogout(this.Ctx, redirectUrl)
}

func (this *CommonController) Login() {
	redirectUrl := this.GetString("redirectUrl")
	sso.RedirectToLogin(this.Ctx, redirectUrl)
}
