// /////////////////////////////////////////////////////////////////////////////
// 从配置中心获取配置

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shima-park/agollo/viper-remote"
	"github.com/spf13/viper"
	"github.com/zpab123/sco"
)

var (
	appConf    *viper.Viper // app 配置
	publicConf *viper.Viper // 公共配置
)

// 设置参数
func setConf() {
	// 设置 viper
	setViper()

	// 服务id
	setMid()

	// 设置集群
	setCluster()
}

// 设置 viper
func setViper() {
	remote.SetAppID(appid)
	p := viper.New()
	p.SetConfigType("prop")

	err := p.AddRemoteProvider("apollo", confUrl[0], "ruixue.bomb")
	if nil != err {
		fmt.Printf("[setViper] 设置 apollo 配置中心失败。err=%s\n", err.Error())
		os.Exit(1)
	}

	err = p.ReadRemoteConfig()
	if nil != err {
		fmt.Printf("[setViper] 从配置中心获取公共配置失败。err=%s\n", err.Error())
		os.Exit(1)
	}

	publicConf = p

	a := viper.New()
	a.SetConfigType("prop")

	err = a.AddRemoteProvider("apollo", confUrl[0], "application")
	if nil != err {
		fmt.Printf("[setViper] 设置 apollo 配置中心失败。err=%s\n", err.Error())
		os.Exit(1)
	}

	err = a.ReadRemoteConfig()
	if nil != err {
		fmt.Printf("[setViper] 读取%s的配置失败。err=%s\n", appid, err.Error())
		os.Exit(1)
	}
	appConf = a
}

// 设置服务id
func setMid() {
	mid := publicConf.GetUint32("mid.match_1v1")
	if mid <= 0 {
		fmt.Println("[setMid] match_1v1 服务id 为0，启动失败。")
		os.Exit(1)
	}

	sco.GetApp().Options.Mid = uint16(mid)
}

// 设置集群
func setCluster() {
	sco.GetApp().Options.Cluster = true

	la := appConf.GetString("RpcServer.Laddr")
	if "" == la {
		fmt.Println("[setCluster] gate 服务 RpcServer.Laddr 为空，启动失败。")
		os.Exit(1)
	}

	sco.GetApp().Options.RpcServer.Laddr = la

	setDis()
}

// 服务发现地址
func setDis() {
	en := publicConf.GetString("Discovery.endpoints")
	if "" == en {
		fmt.Println("[setDis] gate 服务 Discover.endpoints 为空，启动失败。")
		os.Exit(1)
	}

	enl := strings.Split(en, "|")
	if len(enl) <= 0 {
		fmt.Println("[setDis] gate 服务 Discover.endpoints 长度为空，启动失败。")
		os.Exit(1)
	}

	sco.GetApp().Options.Discovery.Endpoints = enl
}
