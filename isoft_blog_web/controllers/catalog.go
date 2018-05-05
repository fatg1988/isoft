package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/common"
	"isoft_blog_web/models"
	"time"
)

type CatalogController struct {
	beego.Controller
}

func (this *CatalogController) List() {
	this.Data["IsCatalogList"] = "IsCatalogList"
	this.Layout = "layout/layout_base.html"
	this.TplName = "catalog/catalog_list.html"
}

func (this *CatalogController) Add() {
	this.Data["IsCatalogAdd"] = "IsCatalogAdd"
	this.Layout = "layout/layout_base.html"
	this.TplName = "catalog/catalog_add.html"
}

func (this *CatalogController) PostAdd() {
	catalog_name := this.GetString("catalog_name")
	catalog_desc := this.GetString("catalog_desc")

	user_name := this.Ctx.Input.Session("UserName").(string)
	catalog := models.Catalog{Author: user_name, CatalogName: catalog_name, CatalogDesc: catalog_desc,
		CreatedBy: user_name, CreatedTime: time.Now(), LastUpdatedBy: user_name, LastUpdatedTime: time.Now()}

	_, err := models.SaveCatalog(&catalog)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
	}
	this.ServeJSON()
}

func (this *CatalogController) PostList() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr["Author"] = this.Ctx.Input.Session("UserName").(string)

	catalogs, count, err := models.QueryCatalog(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	//初始化
	data := make(map[string]interface{}, 1)

	if err == nil {
		data["catalogs"] = catalogs
		data["paginator"] = common.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
	}
	//序列化
	json_obj, err := json.Marshal(data)
	if err == nil {
		this.Data["json"] = string(json_obj)
	} else {
		fmt.Print(err.Error())
	}
	this.ServeJSON()
}
