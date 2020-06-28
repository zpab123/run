// /////////////////////////////////////////////////////////////////////////////
// 从配置中心获取配置

package main

import (
	"fmt"
	"os"

	"github.com/shima-park/agollo/viper-remote"
	"github.com/spf13/viper"
	"github.com/zpab123/sco"
	"github.com/zpab123/sco/app"
)

// 设置参数
func setConf() {
	// 服务id
	setMid()

	// 设置后端
	setBackend()
}

// 设置服务id
func setMid() {
	remote.SetAppID(appid)
	p := viper.New()
	p.SetConfigType("prop")

	err := p.AddRemoteProvider("apollo", confUrl[0], "ruixue.bomb")
	if nil != err {
		fmt.Printf("[setMid] 设置 apollo 配置中心失败。err=%s\n", err.Error())
		os.Exit(1)
	}

	err = p.ReadRemoteConfig()
	if nil != err {
		fmt.Printf("[setMid] 从配置中心获取公共配置失败。err=%s\n", err.Error())
		os.Exit(1)
	}

	mid := p.GetUint32("mid.match_1v1")
	if mid <= 0 {
		fmt.Println("[setMid] match_1v1 服务id 为0，启动失败。")
		os.Exit(1)
	}

	sco.GetApp().Options.Mid = uint16(mid)
}

// 设置后端
func setBackend() {
	sco.GetApp().Options.AppType = app.C_APP_TYPE_BACKEND
	remote.SetAppID(appid)
	v := viper.New()
	v.SetConfigType("prop")

	err := v.AddRemoteProvider("apollo", confUrl[0], "application")
	if nil != err {
		fmt.Printf("[setFrontend] 设置 apollo 配置中心失败。err=%s\n", err.Error())
		os.Exit(1)
	}

	err = v.ReadRemoteConfig()
	if nil != err {
		fmt.Printf("[setFrontend] 读取%s的配置失败。err=%s\n", appid, err.Error())
		os.Exit(1)
	}
}
