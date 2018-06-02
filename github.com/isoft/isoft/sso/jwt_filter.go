package sso

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var (
	isoft_sso_url string
	// 建立一个全局session mananger对象
	globalSessions *session.Manager
)

func init() {
	isoft_sso_url = beego.AppConfig.String("isoft_sso_url")
	// 初始化全局session mananger对象
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

func LoginFilter(ctx *context.Context) {
	if !IsWhiteUrl(ctx) { // 判断是否是白名单中的地址
		successLogin := false // 标记登录信息有效性
		// 从 cookie 中获取 token
		tokenString := ctx.GetCookie("token")
		if tokenString != "" {
			resp, err := http.Get(isoft_sso_url + "/user/checkOrInValidateTokenString?tokenString=" + tokenString + "&operateType=check")
			defer resp.Body.Close()
			if err == nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					jsonStr := string(body)
					var jsonMap map[string]string
					json.Unmarshal([]byte(jsonStr), &jsonMap)
					if jsonMap["status"] == "SUCCESS" {
						successLogin = true
						if ctx.Input.CruSession == nil {
							// 从未访问过是没有 session 的,需要重新创建
							ctx.Input.CruSession, _ = globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
							ctx.Input.CruSession.Set("UserName", jsonMap["username"])
						} else {
							// 登录信息认证通过
							ctx.Input.CruSession.Set("UserName", jsonMap["username"])
						}
					}
				}
			} else {
				successLogin = false
			}
		}
		if !successLogin {
			// 前去登录
			RedirectToLogin(ctx, "")
		}
	}
}

var LoginWhiteList *map[string]string // 登录白名单
var once sync.Once

func GetLoginWhiteList() *map[string]string {
	once.Do(func() {
		m := make(map[string]string)
		m["/common/login"] = "/common/login"
		m["/common/login/"] = "/common/login/"
		m["/common/regist"] = "/common/regist"
		m["/common/regist/"] = "/common/regist/"
		LoginWhiteList = &m
	})
	return LoginWhiteList
}

func IsWhiteUrl(ctx *context.Context) bool {
	fmt.Printf(ctx.Input.URL())
	if _, ok := (*GetLoginWhiteList())[ctx.Input.URL()]; ok {
		return true
	}
	return false
}

func RedirectToLogin(ctx *context.Context, defaultRedirectUrl string) {
	scheme := ctx.Input.Scheme()
	if defaultRedirectUrl != "" {
		defaultRedirectUrl = isoft_sso_url + "/user/login?redirectUrl=" + defaultRedirectUrl
	} else {
		if scheme == "https" {
			defaultRedirectUrl = isoft_sso_url + "/user/login?redirectUrl=" + ctx.Input.Site() + ctx.Input.URI()
		} else {
			defaultRedirectUrl = isoft_sso_url + "/user/login?redirectUrl=" + ctx.Input.Site() + ":" + strconv.Itoa(ctx.Input.Port()) + ctx.Input.URI()
		}
	}
	ctx.Redirect(301, defaultRedirectUrl)
	return
}

func RedirectToLogout(ctx *context.Context, defaultRedirectUrl string) {
	username := ctx.Input.CruSession.Get("UserName").(string)
	// 使 tokenString 失效
	resp, _ := http.Get(isoft_sso_url + "/user/checkOrInValidateTokenString?username=" + username + "&operateType=invalid")
	defer resp.Body.Close()

	// session 失效
	ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)

	// 重新跳往登录页面
	RedirectToLogin(ctx, defaultRedirectUrl)
	return
}
