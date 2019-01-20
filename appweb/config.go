// /////////////////////////////////////////////////////////////////////////////
// 读取配置文件

package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego/config" // 配置文件工具库
	"github.com/astaxie/beego/logs"   // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	mainPath      string              // 程序启动目录
	iniFilePath   = "/conf/conf.conf" // conf.conf 默认路径
	iniAbsPath    string              // conf.conf 绝对路径
	iniconf       config.Configer     // conf.conf 配置文件
	whiteCountry  = make([]string, 0) // 国家白名单
	whiteTimeutc  = make([]string, 0) // 时区白名单
	whiteLanguage = make([]string, 0) // 语言白名单
)

// 初始化
func init() {
	// 获取路径
	mainPath = getMainPath()

	// 绝对路径
	iniAbsPath = filepath.Join(mainPath, iniFilePath)

	// 读取 conf.conf 配置
	readConfIni()

	// 初始化名单信息
	setWhiteCountry()
	setWhiteLanguage()
	setWhiteTimeUtc()
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

// 读取配置文件
func readConfIni() {
	var err error // 错误变量
	iniconf, err = config.NewConfig("ini", iniAbsPath)
	if err != nil {
		logs.Error("conf.conf 配置文件读取错误。err=", err)
	}
}

// 获取 当前 main 程序运行的绝对路径 例如：E:\code\go\go-project\src\test
func getMainPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logs.Error("获取 main 程序 绝对路径失败")
		return ""
	}
	//strings.Replace(dir, "\\", "/", -1)
	return dir
}

// 设置国家白名单
func setWhiteCountry() {
	// 获取字符串
	countryString := iniconf.String("white::country")
	if "" == countryString {
		return
	}

	// 分割字符串
	whiteCountry = strings.Split(countryString, ",")

	// 打印信息
	if false {
		for _, coun := range whiteCountry {
			logs.Info(coun)
		}
	}
}

// 设置时区白名单
func setWhiteTimeUtc() {
	// 获取字符串
	timeString := iniconf.String("white::timeutc")
	if "" == timeString {
		return
	}

	// 分割字符串
	whiteTimeutc = strings.Split(timeString, ",")

	// 打印信息
	if false {
		for _, utc := range whiteTimeutc {
			logs.Info(utc)
		}
	}
}

// 设置语言白名单
func setWhiteLanguage() {
	// 获取字符串
	language := iniconf.String("white::language")
	if "" == language {
		return
	}

	// 分割字符串
	whiteLanguage = strings.Split(language, ",")

	// 打印信息
	if false {
		for _, language := range whiteLanguage {
			logs.Info(language)
		}
	}
}
