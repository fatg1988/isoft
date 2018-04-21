package filter

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"io/ioutil"
	"isoft_sso_tools"
	"net/http"
	"strconv"
	"sync"
)

var isoft_sso_url string

func init() {
	isoft_sso_url = beego.AppConfig.String("isoft_sso_url")
}

func LoginFilter(ctx *context.Context) {
	if !IsWhiteUrl(ctx) { // 判断是否是白名单中的地址
		hasLogin := false
		// 从 cookie 中获取 token
		tokenString := ctx.GetCookie("token")
		if tokenString != "" {
			username, err := isoft_sso_tools.ValidateAndParseJWT(tokenString)
			if err == nil {
				resp, err := http.Get(isoft_sso_url + "/user/checkOrInValidateTokenString?tokenString=" + tokenString + "&operateType=check")
				defer resp.Body.Close()
				if err == nil {
					body, err := ioutil.ReadAll(resp.Body)
					if err == nil {
						jsonStr := string(body)
						var jsonMap map[string]string
						json.Unmarshal([]byte(jsonStr), &jsonMap)
						if jsonMap["status"] == "SUCCESS" {
							hasLogin = true
							// 登录信息认证通过
							ctx.Input.CruSession.Set("UserName", username)
						}
					}
				}
			} else {
				// 登录认证信息不通过
				hasLogin = false
			}
		} else {
			hasLogin = false
		}

		if !hasLogin {
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
