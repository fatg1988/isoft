package models

import "github.com/astaxie/beego/orm"

func SaveBlog(blog *Blog) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(blog)
	return id, err
}

func QueryBlog(condArr map[string]string, page int, offset int) (blogs []Blog, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("blog")
	var cond = orm.NewCondition()

	if value, ok := condArr["Author"]; ok {
		cond = cond.And("Author", value)
	}

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&blogs)
	return
}
