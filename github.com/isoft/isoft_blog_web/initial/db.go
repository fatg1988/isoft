package initial

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/isoft/isoft_blog_web/models"
	"net/url"
)

func InitDB() {

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
	orm.RegisterModel(new(models.Catalog))
	orm.RegisterModel(new(models.Blog))
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
