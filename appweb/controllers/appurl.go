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
	// 获取 get 参数
	this.getOpts()

	// 返回数据
	this.reply()
}

// 获取 get 参数
func (this *AppUrl) getOpts() {
	this.appName = this.GetString("appName")   // app 名字
	this.country = this.GetString("country")   // 国家
	this.language = this.GetString("language") // 语言
	this.timeutc = this.GetString("timeutc")   // 时区

	logs.Info("------------------------------------------------------------------")
	logs.Info("ip=", this.Ctx.Request.RemoteAddr)
	logs.Info("appName=", this.appName)
	logs.Info("country=", this.country)
	logs.Info("timeutc=", this.timeutc)
	logs.Info("language=", this.language)
	logs.Info("------------------------------------------------------------------")
}

// 返回消息
func (this *AppUrl) reply() {
	// 名字不存在
	if "" == this.appName {
		this.Ctx.WriteString("0")
		return
	}

	// 参数检查
	coun := this.CheckCountry()
	utc := this.CheckTimeUtc()
	language := this.CheckLanguage()

	if coun && utc && language {
		webUrl := this.getAppWeb()
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
func (this *AppUrl) CheckCountry() bool {
	if 0 == len(config.WhiteCountry) {
		return false
	}

	for _, coun := range config.WhiteCountry {
		if coun == this.country {
			return true
		}
	}

	return false
}

// 检查时区
func (this *AppUrl) CheckTimeUtc() bool {
	if 0 == len(config.WhiteTimeutc) {
		return false
	}

	for _, utc := range config.WhiteTimeutc {
		if utc == this.timeutc {
			return true
		}
	}

	return false
}

// 检查语言
func (this *AppUrl) CheckLanguage() bool {
	if 0 == len(config.WhiteLanguage) {
		return false
	}

	for _, language := range config.WhiteLanguage {
		if language == this.language {
			return true
		}
	}

	return false
}

// 获取web
func (this *AppUrl) getAppWeb() string {
	// 索引key
	key := fmt.Sprintf("appWeb::%s", this.appName)

	// 参数
	webUrl := config.Iniconf.String(key)

	return webUrl
}
