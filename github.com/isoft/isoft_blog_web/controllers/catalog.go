package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/isoft/isoft/common"
	"github.com/isoft/isoft_blog_web/models"
	"time"
)

type CatalogController struct {
	beego.Controller
}

func (this *CatalogController) More() {
	this.Layout = "layout/layout_front.html"
	this.TplName = "catalog/catalog_more.html"
}

func (this *CatalogController) List() {
	this.Data["IsCatalogList"] = "IsCatalogList"
	this.Layout = "layout/layout_backup.html"
	this.TplName = "catalog/catalog_list.html"
}

func (this *CatalogController) Edit() {
	catalog_id, err := this.GetInt64("catalog_id")
	if err == nil && catalog_id > 0 {
		catalog, err := models.QueryCatalogById(catalog_id)
		if err == nil {
			this.Data["Catalog"] = catalog
		}
	}
	this.Data["IsCatalogAdd"] = "IsCatalogEdit"
	this.Layout = "layout/layout_backup.html"
	this.TplName = "catalog/catalog_edit.html"
}

func (this *CatalogController) PostEdit() {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	user_name := this.Ctx.Input.Session("UserName").(string)

	catalog_id, err := this.GetInt64("catalog_id")
	catalog_name := this.GetString("catalog_name")
	catalog_desc := this.GetString("catalog_desc")
	var catalog models.Catalog
	if err == nil && catalog_id > 0 {
		catalog, err = models.QueryCatalogById(catalog_id)
		if err == nil {
			catalog.CatalogName = catalog_name
			catalog.CatalogDesc = catalog_desc
			catalog.LastUpdatedBy = user_name
			catalog.LastUpdatedTime = time.Now()
		}
	} else {
		catalog = models.Catalog{Author: user_name, CatalogName: catalog_name, CatalogDesc: catalog_desc,
			CreatedBy: user_name, CreatedTime: time.Now(), LastUpdatedBy: user_name, LastUpdatedTime: time.Now()}
	}
	_, err = models.InsertOrUpdateCatalog(&catalog)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}

	this.ServeJSON()
}

func (this *CatalogController) PostList() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页

	// personal="personal"表示查询自己的分类,否则就查询热门分类
	personal := this.GetString("personal")
	if personal == "personal" {
		condArr["Author"] = this.Ctx.Input.Session("UserName").(string)
	} else {
		// 满足热门分类的条件
	}

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

func (this *CatalogController) PostDelete() {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	catalog_id, err := this.GetInt64("catalog_id")
	if err != nil {
		this.ServeJSON()
		return
	}

	err = models.DeleteBlogByCatalogId(catalog_id)
	if err == nil {
		err = models.DeleteCatalogById(catalog_id)
		if err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	}
	this.ServeJSON()
}
