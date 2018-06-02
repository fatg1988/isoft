package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // _ 的作用,并不需要把整个包都导入进来,仅仅是是希望它执行init()函数而已
	"github.com/isoft/isoft_sso_web/models"
	_ "github.com/isoft/isoft_sso_web/routers" // 经常看到Golang代码中出现 _ "controller/home" 类似这种的引用,这里的下划线有什么作用呢?
	// 其实默认每个文件都有一个init函数,加下划线表示引入这个包,仅执行init函数
	"net/url"
)

func initLog() {
	// 控制台输出
	logs.SetLogger(logs.AdapterConsole)
	// 多文件输出
	logs.SetLogger(logs.AdapterMultiFile,
		`{"filename":"../../../../logs/isoft_sso_web.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	// 输出文件名和行号
	logs.EnableFuncCallDepth(true)
	// 异步输出日志
	logs.Async(1e3)
}

func init() {
	// 初始化日志信息
	initLog()

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
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.AppRegister))
	orm.RegisterModel(new(models.LoginRecord))
	orm.RegisterModel(new(models.UserToken))
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
	beego.Run()
}
