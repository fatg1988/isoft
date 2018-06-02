package models

import (
	"github.com/astaxie/beego/orm"
)

func InsertOrUpdateCatalog(catalog *Catalog) (int64, error) {
	o := orm.NewOrm()
	id, err := o.InsertOrUpdate(catalog)
	return id, err
}

func QueryAllCatalog(username string) (catalogs []Catalog, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("catalog").Filter("author", username).All(&catalogs)
	return
}

func QueryCatalog(condArr map[string]string, page int, offset int) (catalogs []Catalog, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	var cond = orm.NewCondition()

	if value, ok := condArr["Author"]; ok {
		cond = cond.And("Author", value)
	}

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&catalogs)
	return
}

func QueryCatalogById(catalog_id int64) (catalog Catalog, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("catalog").Filter("id", catalog_id).One(&catalog)
	return
}

func DeleteCatalogById(catalog_id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("catalog").Filter("id", catalog_id).Delete()
	return
}
