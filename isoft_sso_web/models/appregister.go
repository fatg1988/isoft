package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type AppRegister struct {
	Id              int       `pk json:"id"`
	AppAddress      string    `json:"app_address"`       // 注册的应用地址
	CreatedBy       string    `json:"created_by"`        // 创建人
	CreatedTime     time.Time `json:"created_time"`      // 创建时间
	LastUpdatedBy   string    `json:"last_updated_by"`   // 修改人
	LastUpdatedTime time.Time `json:"last_updated_time"` // 修改时间
}

func AddRegister(appRegister *AppRegister) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(appRegister)
	return id, err
}

// 检查是否经过应用注册
func CheckRegister(app_address string) bool {
	o := orm.NewOrm()
	count, err := o.QueryTable("app_register").Filter("app_address", app_address).Count()
	if err == nil && count > 0 {
		return true
	} else {
		logs.Error("%s has not registe", app_address)
		return false
	}
}

func QueryRegisterCount(appAddress string) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("app_register").Filter("app_address", appAddress).Count()
	return
}

func QueryRegister(condArr map[string]string, page int, offset int) (appregisters []AppRegister, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("app_register")
	var cond = orm.NewCondition()

	if _, ok := condArr["search"]; ok {
		subCond := orm.NewCondition()
		subCond = cond.And("app_address__contains", condArr["search"]).
			Or("created_by", condArr["search"]).Or("last_updated_by", condArr["search"])
		cond = cond.AndCond(subCond)
	}

	qs = qs.OrderBy("-last_updated_time")

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&appregisters)
	return
}
