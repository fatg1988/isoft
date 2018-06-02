package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/isoft/isoft_sso_web/ilearning/util"
	"github.com/isoft/isoft_sso_web/models"
	"time"
)

type AppRegisterController struct {
	beego.Controller
}

func (this *AppRegisterController) AddAppRegister() {
	appAddress := this.GetString("appAddress")
	var appRegister models.AppRegister
	appRegister.AppAddress = appAddress
	appRegister.CreatedBy = "SYSTEM"
	appRegister.LastUpdatedBy = "SYSTEM"
	appRegister.CreatedTime = time.Now()
	appRegister.LastUpdatedTime = time.Now()

	count, err := models.QueryRegisterCount(appAddress)
	if err == nil && count == 0 {
		_, err = models.AddRegister(&appRegister)
		if err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败!"}
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "该系统已注册,请不要重复注册!"}
	}
	this.ServeJSON()
}

func (this *AppRegisterController) AppRegisterList() {
	method := this.Ctx.Request.Method
	if method == "GET" {
		// 跳往管理界面
		this.Layout = "admin/admin_manage_layout.html"
		this.TplName = "admin/appregister_list.html"
	} else {
		condArr := make(map[string]string)
		offset, _ := this.GetInt("offset", 10)            // 每页记录数
		current_page, _ := this.GetInt("current_page", 1) // 当前页

		search := this.GetString("search")
		if search != "" {
			condArr["search"] = search
		}
		appregisters, count, err := models.QueryRegister(condArr, current_page, offset)
		paginator := pagination.SetPaginator(this.Ctx, offset, count)

		if err == nil {
			this.Data["json"] = &map[string]interface{}{"appRegisters": appregisters,
				"paginator": util.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
		}
		this.ServeJSON()
	}
}
