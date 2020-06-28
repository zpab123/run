// /////////////////////////////////////////////////////////////////////////////
// 服务器初始化

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/zpab123/sco"
	"github.com/zpab123/sco/app"
)

var (
	confUrl []string = make([]string, 0) // 配置中心地址
	appid   string   = ""                // app id
)

func init() {
	appid = "match_1v1"
	sco.GetApp().Options.Id = appid

	sco.GetApp().Options.AppType = app.C_APP_TYPE_BACKEND

	// 解析启动参数
	parseArgs()

	// 设置参数
	setConf()
}

// 解析启动参数
func parseArgs() {
	ccu := flag.String("ccu", "", "config center url") // 配置中心地址集合

	// 解析参数
	flag.Parse()

	// 配置中心
	if *ccu == "" {
		fmt.Printf("配置中心参数 [ccu] 为空，启动失败！")
		os.Exit(1)
	} else {
		parseConfUrl(*ccu)
	}
}

// 解析配置中心地址
func parseConfUrl(u string) {
	us := strings.Split(u, "|")
	if len(us) <= 0 {
		fmt.Printf("配置中心参数 [ccu] 错误，启动失败！")
		os.Exit(1)
		return
	}

	confUrl = append(confUrl, us...)
}
