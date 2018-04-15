package controllers

import (
	"isoft_sso_web/ilearning/util"
	"isoft_sso_web/models"
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"strings"
	"time"
)

var originConfig string
var aes_sso_key string

var globalSessions *session.Manager

func init() {
	originConfig = beego.AppConfig.String("origin")
	aes_sso_key = beego.AppConfig.String("aes_sso_key")

	// 初始化一个全局的变量用来存储 session 控制器
	sessionConfig := &session.ManagerConfig{
		CookieName:      "beegosessionid", // 客户端存储 cookie 的名字
		EnableSetCookie: true,
		Gclifetime:      3600, // 触发 GC 的时间
		Maxlifetime:     3600, // 服务器端存储的数据的过期时间
		Secure:          false,
		CookieLifeTime:  3600,
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()

}

type UserController struct {
	beego.Controller
}

func (this *UserController) DeleteToken() {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": "Token 删除失败!"}
	// 获取密文 token
	token := this.GetString("token")
	// 对密文 token 进行解密
	key := []byte(aes_sso_key)
	b, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		// 自定义解码方式,主要供非 go 代码方式调用
		token = strings.Replace(token, "BAAAAAAAAB", "+", -1)
		token = strings.Replace(token, "ABBBBBBBBA", "/", -1)
		b, err = base64.StdEncoding.DecodeString(token)
	}
	origData, err := util.AesDecrypt(b, key)
	if err == nil {
		tokenStr := string(origData)
		var tokenMap map[string]string
		json.Unmarshal([]byte(tokenStr), &tokenMap)
		ssoSessionId := tokenMap["ssoSessionId"]
		sess, _ := globalSessions.GetSessionStore(ssoSessionId)
		// session 释放
		sess.SessionRelease(this.Ctx.ResponseWriter)
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "msg": "Token 删除成功!"}
	}
	this.ServeJSON()
}

func (this *UserController) Logout() {
	this.DelSession("username")
	this.Redirect("/user/login", 302)
}

func (this *UserController) CheckLogin() {
	this.Data["json"] = &map[string]interface{}{"userName": "", "status": "ERROR", "msg": "认证失败!"}

	// 获取密文 token
	token := this.GetString("token")
	// 对密文 token 进行解密
	key := []byte(aes_sso_key)
	b, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		token = strings.Replace(token, "BAAAAAAAAB", "+", -1)
		token = strings.Replace(token, "ABBBBBBBBA", "/", -1)
		b, err = base64.StdEncoding.DecodeString(token)
	}
	origData, err := util.AesDecrypt(b, key)
	if err == nil {
		tokenStr := string(origData)
		var tokenMap map[string]string
		json.Unmarshal([]byte(tokenStr), &tokenMap)

		ssoSessionId := tokenMap["ssoSessionId"]
		userName := tokenMap["userName"]
		isLogin := tokenMap["isLogin"]

		sess, _ := globalSessions.GetSessionStore(ssoSessionId)
		if sess.Get("userName") == userName && sess.Get("isLogin") == isLogin {
			this.Data["json"] = &map[string]interface{}{"userName": tokenMap["userName"], "isLogin": tokenMap["isLogin"], "status": "SUCCESS", "msg": "认证成功!"}
		} else {
			// session 释放
			sess.SessionRelease(this.Ctx.ResponseWriter)
		}
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
	// 将用户登录信息添加到 session 中去
	this.SetSession("UserName", user.UserName)
	sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	sess.Set("ssoSessionId", sess.SessionID())
	sess.Set("userName", user.UserName)
	sess.Set("isLogin", "isLogin")
	// 设置 cookie 信息
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	token := make(map[string]string, 10)
	token["ssoSessionId"] = sess.SessionID()
	token["userName"] = username
	token["isLogin"] = "isLogin"
	b, er := json.Marshal(token)
	if er == nil {
		// 使用 AES 加密算法进行加密
		key := []byte(aes_sso_key)
		result, err := util.AesEncrypt([]byte(string(b)), key)
		if err == nil {
			// 加密成功后将密文 token 添加到 cookie 中
			this.Ctx.SetCookie("token", base64.StdEncoding.EncodeToString(result), 100, "/")
		}
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

func CheckOrigin(origin string) bool {
	if origin == originConfig {
		return true
	}
	return false
}
