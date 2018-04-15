package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"time"
)

type User struct {
	Id       				int 		`pk json:"id"`
	UserName 				string 		`json:"user_name"`
	PassWd   				string 		`json:"pass_wd"`
	CreatedBy				string 		`json:"created_by"`			// 创建人
	CreatedTime				time.Time	`json:"created_time"`		// 创建时间
	LastUpdatedBy			string		`json:"last_updated_by"`	// 修改人
	LastUpdatedTime			time.Time	`json:"last_updated_time"`	// 修改时间
}

func SaveUser(user User) error {
	o := orm.NewOrm()
	count, _ := o.QueryTable("user").Filter("user_name",user.UserName).Count()
	if count > 0{
		return errors.New("用户已注册!")
	}else{
		_, err := o.Insert(&user)
		return err
	}
	return nil
}

func QueryUser(username,passwd string) (user User, err error)  {
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("user_name",username).Filter("pass_wd",passwd).One(&user)
	return
}


