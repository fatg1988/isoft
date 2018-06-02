package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

func InsertOrUpdateBlog(blog *Blog) (id int64, err error) {
	o := orm.NewOrm()
	if blog.Id > 0 {
		id, err = o.Update(blog)
	} else {
		id, err = o.Insert(blog)
	}
	return
}

func QueryBlog(condArr map[string]interface{}, page int, offset int) (blogs []Blog, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("blog")
	var cond = orm.NewCondition()

	if value, ok := condArr["Author"]; ok {
		cond = cond.And("Author", value)
	}

	if catalog_id, ok := condArr["catalog_id"]; ok {
		cond = cond.And("Catalog__Id", catalog_id)
	}

	if BlogStatus, ok := condArr["BlogStatus"]; ok {
		cond = cond.And("BlogStatus", BlogStatus)
	}

	if search_text, ok := condArr["search_text"]; ok { // 根据博客分类/博文名称/搜索关键词/作者等信息查询
		var search_cond = orm.NewCondition()
		search_cond = search_cond.Or("Catalog__CatalogName__contains", search_text).Or("BlogTitle__contains", search_text).
			Or("KeyWords__contains", search_text).Or("Author__contains", search_text)
		cond = cond.AndCond(search_cond)
	}

	qs = qs.SetCond(cond)

	if _, ok := condArr["querysOrder"]; ok {
		querysOrder := condArr["querysOrder"].(string)
		// 多个排序条件使用 @ 符号进行分割
		querysOrderArr := strings.Split(querysOrder, "@")
		for _, v := range querysOrderArr {
			qs = qs.OrderBy(v)
		}
	}

	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&blogs)
	return
}

func DeleteBlogByCatalogId(catalog_id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("blog").Filter("catalog_id", catalog_id).Delete()
	return
}

func DeleteBlogById(blog_id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("blog").Filter("id", blog_id).Delete()
	return
}

func QueryBlogById(blog_id int64) (blog Blog, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("blog").Filter("id", blog_id).One(&blog)
	return
}

func PublishBlogById(blog_id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("blog").Filter("id", blog_id).Update(orm.Params{"BlogStatus": 1})
	return
}
