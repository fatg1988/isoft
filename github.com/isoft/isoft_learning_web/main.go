package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // _ 的作用,并不需要把整个包都导入进来,仅仅是是希望它执行init()函数而已
	"github.com/isoft/isoft/sso"
	"github.com/isoft/isoft_learning_web/models"
	_ "github.com/isoft/isoft_learning_web/routers"
	"net/url"
)

func init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.pass")
	timezone := beego.AppConfig.String("db.timezone")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.SetMaxIdleConns("default", 1000) // SetMaxIdleConns用于设置闲置的连接数
	orm.SetMaxOpenConns("default", 2000) // SetMaxOpenConns用于设置最大打开的连接数,默认值为0表示不限制

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	registerModel()

	createTable() // 开启自动建表
}

func registerModel() {
	orm.RegisterModel(new(models.Course))
	orm.RegisterModel(new(models.CourseVedio))
	orm.RegisterModel(new(models.Favorite))
	orm.RegisterModel(new(models.TopicTheme))
	orm.RegisterModel(new(models.TopicReply))
	orm.RegisterModel(new(models.Note))
	orm.RegisterModel(new(models.Configuration))
}

// 自动建表
func createTable() {
	name := "default"                          // 数据库别名
	force := false                             // 不强制建数据库
	verbose := true                            // 打印建表过程
	err := orm.RunSyncdb(name, force, verbose) // 建表
	if err != nil {
		beego.Error(err)
	}
}

func main() {
	beego.InsertFilter("/*", beego.BeforeExec, sso.LoginFilter)
	beego.Run()
}
