// /////////////////////////////////////////////////////////////////////////////
// gate 服务器

package main

import (
	"flag"
	"log"
	"run/center/arith"

	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr       = flag.String("addr", "localhost:8972", "server address")
	consulAddr = flag.String("consulAddr", "localhost:8500", "consul address")
	basePath   = flag.String("base", "/arith/arith", "prefix path")
)

// 主入口
func main() {
	// 解析参数
	flag.Parse()

	// 创建服务器
	s := server.NewServer()

	// 添加插件
	addRegistryPlugin(s)

	// 注册服务
	s.RegisterName("Arith", new(arith.Arith), "")

	// 开启服务
	s.Serve("tcp", *addr)
}

// 添加插件
func addRegistryPlugin(s *server.Server) {
	// 创建插件
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ConsulServers:  []string{*consulAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}

	// 插件开始
	err := r.Start()

	if err != nil {
		log.Fatal(err)
	}

	// 添加插件
	s.Plugins.Add(r)
}
