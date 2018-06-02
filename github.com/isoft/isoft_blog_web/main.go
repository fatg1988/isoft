package main

import (
	"github.com/astaxie/beego"
	"github.com/isoft/isoft/sso"
	"github.com/isoft/isoft_blog_web/initial"
	_ "github.com/isoft/isoft_blog_web/routers"
)

func init() {
	initial.InitLog()
	initial.InitDB()
}

func main() {
	// 进行相关组件注册
	StartRegister()
	beego.Run()
}

func StartRegister() {
	// 登录过滤器
	beego.InsertFilter("/*", beego.BeforeExec, sso.LoginFilter)
}
