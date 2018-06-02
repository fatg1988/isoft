package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/isoft/isoft_sso_web/models"
	"strings"
	"time"
)

var origin_list string

func init() {
	origin_list = beego.AppConfig.String("origin_list")
}

type UserController struct {
	beego.Controller
}

func (this *UserController) CheckOrInValidateTokenString() {
	tokenString := this.GetString("tokenString")
	username := this.GetString("username")
	operateType := this.GetString("operateType")
	if operateType == "check" {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
		username, err := ValidateAndParseJWT(tokenString)
		if err == nil {
			_, err = models.QueryUserToken(username)
			if err == nil {
				this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "username": username}
			}
		}
	} else {
		// 删除 tokenString,使客户端登录信息失效
		userToken, _ := models.QueryUserToken(username)
		models.DeleteUserToken(userToken)
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *UserController) Regist() {
	Method := this.Ctx.Request.Method
	if Method == "GET" {
		this.TplName = "regist.html"
	} else {
		var user models.User
		inputs := this.Input()
		user.UserName = inputs.Get("username")
		user.PassWd = inputs.Get("passwd")
		user.CreatedBy = "SYSTEM"
		user.CreatedTime = time.Now()
		user.LastUpdatedBy = "SYSTEM"
		user.LastUpdatedTime = time.Now()
		err := models.SaveUser(user)

		//初始化
		data := make(map[string]interface{}, 1)

		if err == nil {
			data["status"] = "SUCCESS"
		} else {
			data["Status"] = "ERROR"
			data["ErrorCode"] = err.Error()
			data["ErrorMsg"] = err.Error()
		}
		//序列化
		json_obj, err := json.Marshal(data)
		if err == nil {
			this.Data["json"] = string(json_obj)
		}
		this.ServeJSON()
	}
}

func (this *UserController) Login() {
	if this.Ctx.Request.Method == "GET" {
		this.TplName = "login.html"
	} else {
		// referer显示来源页面的完整地址,而origin显示来源页面的origin: protocal+host,不包含路径等信息,也就不会包含含有用户信息的敏感内容
		// referer存在于所有请求,而origin只存在于post请求,随便在页面上点击一个链接将不会发送origin
		// 因此origin较referer更安全,多用于防范CSRF攻击
		referer := this.Ctx.Input.Referer()
		origin := this.Ctx.Request.Header.Get("origin")
		username := this.Input().Get("username")
		passwd := this.Input().Get("passwd")
		if IsAdminUser(username) { // 是管理面账号
			AdminUserLogin(origin, this, username, referer)
		} else {
			CommonUserLogin(referer, origin, username, passwd, this)
		}
	}
}
func CommonUserLogin(referer string, origin string, username string, passwd string, this *UserController) {
	referers := strings.Split(referer, "/user/login?redirectUrl=")
	if CheckOrigin(origin) && len(referers) == 2 && CheckOrigin(referers[0]) && IsValidRedirectUrl(referers[1]) {
		user, err := models.QueryUser(username, passwd)
		if err == nil && &user != nil {
			SuccessedLogin(username, this, origin, referer, user, referers)
		} else {
			ErrorAccountLogin(username, this, origin, referer)
		}
	} else {
		ErrorAuthorizedLogin(username, this, origin, referer)
	}
}

func SuccessedLogin(username string, this *UserController, origin string, referer string, user models.User, referers []string) {
	var loginLog models.LoginRecord
	loginLog.UserName = username
	loginLog.LoginIp = this.Ctx.Input.IP()
	loginLog.Origin = origin
	loginLog.Referer = referer
	loginLog.LoginStatus = "success"
	loginLog.LoginResult = "SUCCESS"
	loginLog.CreatedBy = "SYSTEM"
	loginLog.CreatedTime = time.Now()
	loginLog.LastUpdatedBy = "SYSTEM"
	loginLog.LastUpdatedTime = time.Now()
	models.AddLoginRecord(loginLog)

	// 设置 cookie 信息
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")

	tokenString, err := CreateJWT(username)
	if err == nil {

		var userToken models.UserToken
		userToken.UserName = username
		userToken.TokenString = tokenString
		userToken.CreatedBy = "SYSTEM"
		userToken.CreatedTime = time.Now()
		userToken.LastUpdatedBy = "SYSTEM"
		userToken.LastUpdatedTime = time.Now()
		models.SaveUserToken(userToken)

		// 设置 cookie 有效时间为 24 小时
		this.Ctx.SetCookie("token", tokenString, 3600*24, "/")
	}

	// 则重定向到 redirectUrl,原生的方法是：w.Header().Set("Location", "http://www.baidu.com") w.WriteHeader(301)
	this.Redirect(referers[1], 301)
}

func ErrorAuthorizedLogin(username string, this *UserController, origin string, referer string) {
	var loginLog models.LoginRecord
	loginLog.UserName = username
	loginLog.LoginIp = this.Ctx.Input.IP()
	loginLog.Origin = origin
	loginLog.Referer = referer
	if !CheckOrigin(origin) {
		loginLog.LoginStatus = "origin_error"
	} else {
		loginLog.LoginStatus = "refer_error"
	}
	loginLog.LoginResult = "FAILED"
	loginLog.CreatedBy = "SYSTEM"
	loginLog.CreatedTime = time.Now()
	loginLog.LastUpdatedBy = "SYSTEM"
	loginLog.LastUpdatedTime = time.Now()
	models.AddLoginRecord(loginLog)
	this.TplName = "403.html"
}

func ErrorAccountLogin(username string, this *UserController, origin string, referer string) {
	var loginLog models.LoginRecord
	loginLog.UserName = username
	loginLog.LoginIp = this.Ctx.Input.IP()
	loginLog.Origin = origin
	loginLog.Referer = referer
	loginLog.LoginStatus = "account_error"
	loginLog.LoginResult = "FAILED"
	loginLog.CreatedBy = "SYSTEM"
	loginLog.CreatedTime = time.Now()
	loginLog.LastUpdatedBy = "SYSTEM"
	loginLog.LastUpdatedTime = time.Now()
	models.AddLoginRecord(loginLog)
	this.Data["ErrorMsg"] = "用户名或密码不正确!"
	this.TplName = "login.html"
}

func AdminUserLogin(origin string, this *UserController, username string, referer string) {
	if CheckOrigin(origin) { // 非跨站点
		// 跳往管理界面
		this.Layout = "admin/admin_manage_layout.html"
		this.TplName = "admin/admin_manage_default.html"
	} else {
		ErrorAuthorizedLogin(username, this, origin, referer)
	}
}

func IsValidRedirectUrl(redirectUrl string) bool {
	if redirectUrl != "" && IsHttpProtocol(redirectUrl) {
		// 截取协议名称
		arr := strings.Split(redirectUrl, "//")
		protocol := arr[0]
		// 截取域名
		a1 := arr[1]
		host := strings.Split(a1, "/")[0]
		return CheckRegister(protocol + "//" + host)
	} else {
		return false
	}
}

func IsAdminUser(user_name string) bool {
	if user_name == "admin1" {
		return true
	}
	return false
}

func IsHttpProtocol(url string) bool {
	if strings.HasPrefix(url, "http") || strings.HasPrefix(url, "https") {
		return true
	}
	return false
}

// 判断是否经过注册
func CheckRegister(registUrl string) bool {
	return models.CheckRegister(registUrl)
}

// 验证 origin 是否合法
func CheckOrigin(origin string) bool {
	origin_slice := strings.Split(origin_list, ",")
	for _, _origin := range origin_slice {
		if origin == _origin {
			return true
		}
	}
	logs.Warn("origin error for %s", origin)
	return false
}
