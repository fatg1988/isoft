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

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Detail() {
	blog_id, err := this.GetInt64("blog_id")
	detail := this.GetString("detail")
	if err == nil {
		if detail == "detail" {
			blog, err := models.QueryBlogById(blog_id)
			if err == nil {
				this.Data["BlogContent"] = blog.Content
				this.TplName = "blog/_blog_detail.html"
			}
		} else {
			this.Data["BlogId"] = blog_id
			this.Layout = "layout/layout_front.html"
			this.TplName = "blog/blog_detail.html"
		}
	}
}

func (this *BlogController) Search() {
	this.Layout = "layout/layout_front.html"
	this.TplName = "blog/blog_search.html"
}

func (this *BlogController) List() {
	this.Data["IsBlogList"] = "IsBlogList"
	this.Layout = "layout/layout_backup.html"
	this.TplName = "blog/blog_list.html"
}

func (this *BlogController) Edit() {
	user_name := this.Ctx.Input.Session("UserName").(string)
	blog_id, err := this.GetInt64("blog_id")
	if err == nil && blog_id > 0 {
		blog, err := models.QueryBlogById(blog_id)
		if err == nil {
			this.Data["Blog"] = blog
		}
	}
	catalogs, err := models.QueryAllCatalog(user_name)
	if err == nil {
		this.Data["Catalogs"] = &catalogs
	}
	this.Data["IsBlogAdd"] = "IsBlogEdit"
	this.Layout = "layout/layout_backup.html"
	this.TplName = "blog/blog_edit.html"
}

func (this *BlogController) PostEdit() {
	blog_id, err := this.GetInt64("blog_id")
	blog_title := this.GetString("blog_title")
	key_words := this.GetString("key_words")
	catalog_id, _ := this.GetInt64("catalog_id", -1)
	blog_type, _ := this.GetInt8("blog_type", -1)
	blog_status, _ := this.GetInt8("blog_status", -1)
	content := this.GetString("content")
	user_name := this.Ctx.Input.Session("UserName").(string)
	catalog, _ := models.QueryCatalogById(catalog_id)
	var blog models.Blog
	if err == nil && catalog_id > 0 {
		blog, err = models.QueryBlogById(blog_id)
		if err == nil {
			blog.BlogTitle = blog_title
			blog.KeyWords = key_words
			blog.BlogType = blog_type
			blog.BlogStatus = blog_status
			blog.Catalog = &catalog
			blog.Content = content
			blog.Edits = blog.Edits + 1
			blog.LastUpdatedBy = user_name
			blog.LastUpdatedTime = time.Now()
		}
	} else {
		blog = models.Blog{
			Author:          user_name,
			BlogTitle:       blog_title,
			KeyWords:        key_words,
			Catalog:         &catalog,
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
	}
	_, err = models.InsertOrUpdateBlog(&blog)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
	}
	this.ServeJSON()
}

func (this *BlogController) PostList() {
	condArr := make(map[string]interface{})
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	catalog_id, _ := this.GetInt64("catalog_id", -1)
	search_text := this.GetString("search_text")

	// personal="personal"表示查询自己的博文,否则就查询热门博文
	personal := this.GetString("personal")
	if personal == "personal" {
		condArr["Author"] = this.Ctx.Input.Session("UserName").(string)
	} else {
		// 满足热门博文的条件,默认按照浏览次数排行
		condArr["querysOrder"] = "-Views"
		// 默认查询已发布的博文
		condArr["BlogStatus"] = 1
	}

	if catalog_id > 0 {
		condArr["catalog_id"] = catalog_id
	}

	if search_text != "" {
		condArr["search_text"] = search_text
	}

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

func (this *BlogController) PostDelete() {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	blog_id, err := this.GetInt64("blog_id")
	if err != nil {
		this.ServeJSON()
		return
	}
	err = models.DeleteBlogById(blog_id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *BlogController) PostPublish() {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	blog_id, err := this.GetInt64("blog_id")
	if err != nil {
		this.ServeJSON()
		return
	}
	err = models.PublishBlogById(blog_id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}
