// /////////////////////////////////////////////////////////////////////////////
// 读取配置文件

package config

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	WhiteCountry  = make([]string, 0) // 国家白名单
	WhiteTimeutc  = make([]string, 0) // 时区白名单
	WhiteLanguage = make([]string, 0) // 语言白名单
)

// 初始化
func init() {

	// 读取国家白名单
	readWhiteCountry()
	// 读取时区白名单
	readWhiteTimeUtc()
	// 读取语言白名单
	readWhiteLanguage()
}

// /////////////////////////////////////////////////////////////////////////////
// public api

// /////////////////////////////////////////////////////////////////////////////
// 对象

// 白名单配置
type WhiteConf struct {
	Country  []string // 国家切片
	Timeutc  []int    // 时区切片
	Language []string // 系统语言切片
}

// 黑名单配置
type BlackConf struct {
	Country  []string // 国家切片
	Timeutc  []int    // 时区切片
	Language []string // 系统语言切片
}

// /////////////////////////////////////////////////////////////////////////////
// private api

// 读取国家白名单
func readWhiteCountry() {
	c := beego.AppConfig.String("white::country")
	if "" == c {
		logs.Error("读取国家白名单出错")
		return
	}

	// 分割字符串
	WhiteCountry = strings.Split(c, ",")

	// 打印信息
	logs.Info("国家白名单如下：")
	for _, coun := range WhiteCountry {
		logs.Info(coun)
	}
}

// 读取时区白名单
func readWhiteTimeUtc() {
	// 获取字符串
	ts := beego.AppConfig.String("white::timeutc")
	if "" == ts {
		logs.Error("读取时区白名单出错")
		return
	}

	// 分割字符串
	WhiteTimeutc = strings.Split(ts, ",")

	// 打印信息
	logs.Info("时区白名单如下：")
	for _, utc := range WhiteTimeutc {
		logs.Info(utc)
	}
}

// 设置语言白名单
func readWhiteLanguage() {
	ln := beego.AppConfig.String("white::language")
	if "" == ln {
		logs.Error("设置语言白名单出错")
		return
	}

	// 分割字符串
	WhiteLanguage = strings.Split(ln, ",")

	// 打印信息
	logs.Info("语言白名单如下：")
	for _, language := range WhiteLanguage {
		logs.Info(language)
	}
}
