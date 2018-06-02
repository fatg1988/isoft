package initial

import (
	"github.com/astaxie/beego/logs"
)

func InitLog() {

	// 控制台输出
	logs.SetLogger(logs.AdapterConsole)
	// 多文件输出
	logs.SetLogger(logs.AdapterMultiFile,
		`{"filename":"../../../../logs/isoft_blog_web.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	// 输出文件名和行号
	logs.EnableFuncCallDepth(true)
	// 异步输出日志
	logs.Async(1e3)

}
