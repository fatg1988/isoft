package main

import (
	"github.com/astaxie/beego"
	"isoft/sso"
	"isoft_blog_web/initial"
	_ "isoft_blog_web/routers"
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
