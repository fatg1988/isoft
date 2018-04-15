package filter

import (
	"github.com/astaxie/beego/context"
	"sync"
	"github.com/astaxie/beego"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"fmt"
	"strconv"
)

var isoft_sso_url string

func init()  {
	isoft_sso_url = beego.AppConfig.String("isoft_sso_url")
}

func LoginFilter(ctx *context.Context) {
	if !IsWhiteUrl(ctx){				// 判断是否是白名单中的地址
		var hasLogin bool
		// 从 cookie 中获取 token
		token := ctx.GetCookie("token")
		if token != ""{
			resp, err := http.Get(isoft_sso_url + "/user/checkLogin?token=" + url.QueryEscape(token))
			defer resp.Body.Close()
			if err == nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					jsonStr := string(body)
					var jsonMap map[string]string
					json.Unmarshal([]byte(jsonStr), &jsonMap)
					if jsonMap["status"] == "SUCCESS"{
						ctx.Input.CruSession.Set("UserName", jsonMap["userName"])
						ctx.Input.CruSession.Set("isLogin", jsonMap["isLogin"])
						hasLogin = true
					}else{
						hasLogin = false
					}
				}
			}
		}else{
			hasLogin = false
		}

		if !hasLogin{
			// 前去登录
			RedirectToLogin(ctx, "")
		}
	}
}

var LoginWhiteList *map[string]string				// 登录白名单
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
	if _,ok := (*GetLoginWhiteList())[ctx.Input.URL()]; ok{
		return true
	}
	return false
}

func RedirectToLogin(ctx *context.Context, defaultRedirectUrl string){
	scheme := ctx.Input.Scheme()
	if defaultRedirectUrl != ""{
		defaultRedirectUrl = isoft_sso_url + "/user/login?redirectUrl=" + defaultRedirectUrl
	}else{
		if scheme == "https"{
			defaultRedirectUrl = isoft_sso_url + "/user/login?redirectUrl=" + ctx.Input.Site() + ctx.Input.URI()
		}else{
			defaultRedirectUrl = isoft_sso_url + "/user/login?redirectUrl=" + ctx.Input.Site() + ":" + strconv.Itoa(ctx.Input.Port()) + ctx.Input.URI()
		}
	}
	ctx.Redirect(301, defaultRedirectUrl)
	return
}

func RedirectToLogout(ctx *context.Context, defaultRedirectUrl string){
	// session 失效
	ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
	// sso session失效
	token := ctx.GetCookie("token") // 从 cookie 中获取 token
	http.Get(isoft_sso_url + "/user/deleteToken?token=" + url.QueryEscape(token))
	// 重新跳往登录页面
	RedirectToLogin(ctx, defaultRedirectUrl)
	return
}