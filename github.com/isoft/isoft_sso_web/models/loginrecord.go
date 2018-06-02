package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type LoginRecord struct {
	Id              int       `pk json:"id"`
	UserName        string    `json:"user_name"` // 登录用户名
	LoginIp         string    `json:"login_ip"`  // 登录用户IP
	Origin          string    `json:"origin"`
	Referer         string    `json:"referer"`
	LoginStatus     string    `json:"login_status"`      // 登录状态
	LoginResult     string    `json:"login_result"`      // 登录结果
	CreatedBy       string    `json:"created_by"`        // 创建人
	CreatedTime     time.Time `json:"created_time"`      // 创建时间
	LastUpdatedBy   string    `json:"last_updated_by"`   // 修改人
	LastUpdatedTime time.Time `json:"last_updated_time"` // 修改时间
}

func AddLoginRecord(log LoginRecord) error {
	o := orm.NewOrm()
	_, err := o.Insert(&log)
	return err
}

func QueryLoginRecord(condArr map[string]string, page int, offset int) (loginRecords []LoginRecord, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("login_record")
	var cond = orm.NewCondition()

	if _, ok := condArr["search"]; ok {
		subCond := orm.NewCondition()
		subCond = cond.And("login_ip__contains", condArr["search"]).
			Or("origin", condArr["search"]).Or("refer", condArr["search"])
		cond = cond.AndCond(subCond)
	}

	qs = qs.OrderBy("-last_updated_time")

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&loginRecords)
	return
}
