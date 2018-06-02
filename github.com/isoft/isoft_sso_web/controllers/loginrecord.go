package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/isoft/isoft_sso_web/ilearning/util"
	"github.com/isoft/isoft_sso_web/models"
)

type LoginRecordController struct {
	beego.Controller
}

func (this *LoginRecordController) LoginRecordList() {
	method := this.Ctx.Request.Method
	if method == "GET" {
		// 跳往管理界面
		this.Layout = "admin/admin_manage_layout.html"
		this.TplName = "admin/loginrecord_list.html"
	} else {
		condArr := make(map[string]string)
		offset, _ := this.GetInt("offset", 10)            // 每页记录数
		current_page, _ := this.GetInt("current_page", 1) // 当前页

		search := this.GetString("search")
		if search != "" {
			condArr["search"] = search
		}
		loginrecords, count, err := models.QueryLoginRecord(condArr, current_page, offset)
		paginator := pagination.SetPaginator(this.Ctx, offset, count)

		if err == nil {
			this.Data["json"] = &map[string]interface{}{"loginRecords": loginrecords,
				"paginator": util.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
		}
		this.ServeJSON()
	}
}
