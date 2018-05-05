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

type BlogController struct {
	beego.Controller
}

func (this *BlogController) List() {
	this.Data["IsBlogList"] = "IsBlogList"
	this.Layout = "layout/layout_base.html"
	this.TplName = "blog/blog_list.html"
}

func (this *BlogController) Add() {
	user_name := this.Ctx.Input.Session("UserName").(string)
	catalogs, err := models.QueryAllCatalog(user_name)
	if err == nil {
		this.Data["Catalogs"] = &catalogs
	}
	this.Data["IsBlogAdd"] = "IsBlogAdd"
	this.Layout = "layout/layout_base.html"
	this.TplName = "blog/blog_add.html"
}

func (this *BlogController) PostAdd() {
	blog_title := this.GetString("blog_title")
	key_words := this.GetString("key_words")
	catalog_id, err := this.GetInt64("catalog_id", -1)
	blog_type, err := this.GetInt8("blog_type", -1)
	blog_status, err := this.GetInt8("blog_status", -1)
	content := this.GetString("content")
	user_name := this.Ctx.Input.Session("UserName").(string)

	blog := models.Blog{
		Author:          user_name,
		BlogTitle:       blog_title,
		KeyWords:        key_words,
		CatalogId:       catalog_id,
		Content:         content,
		BlogType:        blog_type,
		BlogStatus:      blog_status,
		Views:           0,
		Edits:           1,
		CreatedBy:       user_name,
		CreatedTime:     time.Now(),
		LastUpdatedBy:   user_name,
		LastUpdatedTime: time.Now(),
	}

	_, err = models.SaveBlog(&blog)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
	}
	this.ServeJSON()
}

func (this *BlogController) PostList() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr["Author"] = this.Ctx.Input.Session("UserName").(string)

	blogs, count, err := models.QueryBlog(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	//初始化
	data := make(map[string]interface{}, 1)

	if err == nil {
		data["blogs"] = blogs
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
