// /////////////////////////////////////////////////////////////////////////////
// 服务器初始化

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	confUrl []string = make([]string, 0) // 配置中心地址
	appid   string   = ""                // app id
)

func init() {
	// 解析启动参数
	parseArgs()
}

// 解析启动参数
func parseArgs() {
	id := flag.String("id", "", "app id")              // app 唯一id标识
	ccu := flag.String("ccu", "", "config center url") // 配置中心地址集合

	// 解析参数
	flag.Parse()

	// id 解析
	if *id == "" {
		fmt.Printf("appid参数 [id] 为空，启动失败！")
		os.Exit(1)
	} else {
		appid = *id
	}

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
