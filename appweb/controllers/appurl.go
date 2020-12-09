package controllers

import (
	"appweb/config"
	"fmt"

	"github.com/astaxie/beego"      //  引擎
	"github.com/astaxie/beego/logs" // 日志库
)

// 请求需要显示的网页地址
type AppUrl struct {
	beego.Controller
	appName  string // 国家
	country  string // 语言
	timeutc  string // 时区
	language string // 语言
}

// get 方法
func (this *AppUrl) Get() {
	appName := this.GetString("aN")    // app 名字
	country := this.GetString("count") // 国家
	language := this.GetString("lan")  // 语言
	timeutc := this.GetString("tu")    // 时区

	logs.Info("------------------------------------------------------------------")
	logs.Info("ip=", this.Ctx.Request.RemoteAddr)
	logs.Info("appName=", appName)
	logs.Info("country=", country)
	logs.Info("timeutc=", timeutc)
	logs.Info("language=", language)
	logs.Info("------------------------------------------------------------------")

	// 名字不存在
	if "" == appName {
		this.Ctx.WriteString("0")
		return
	}

	// 参数检查
	coun := this.CheckCountry(country)
	utc := this.CheckTimeUtc(timeutc)
	lan := this.CheckLanguage(language)

	if coun && utc && lan {
		webUrl := this.getAppWeb(appName)
		if "" == webUrl {
			this.Ctx.WriteString("0")
		} else {
			this.Ctx.WriteString(webUrl)
		}
	} else {
		this.Ctx.WriteString("0")
	}

}

// 检查国家
func (this *AppUrl) CheckCountry(country string) bool {
	if 0 == len(config.WhiteCountry) {
		return false
	}

	for _, coun := range config.WhiteCountry {
		if coun == country {
			return true
		}
	}

	return false
}

// 检查时区
func (this *AppUrl) CheckTimeUtc(tuc string) bool {
	if 0 == len(config.WhiteTimeutc) {
		return false
	}

	for _, utc := range config.WhiteTimeutc {
		if utc == tuc {
			return true
		}
	}

	return false
}

// 检查语言
func (this *AppUrl) CheckLanguage(lan string) bool {
	if 0 == len(config.WhiteLanguage) {
		return false
	}

	for _, language := range config.WhiteLanguage {
		if language == lan {
			return true
		}
	}

	return false
}

// 获取web
func (this *AppUrl) getAppWeb(name string) string {
	// 索引key
	key := fmt.Sprintf("appWeb::%s", name)

	// 参数
	webUrl := beego.AppConfig.String(key)

	return webUrl
}
