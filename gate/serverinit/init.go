// /////////////////////////////////////////////////////////////////////////////
// 服务器初始化

package serverinit

import (
	"github.com/zpab123/sco"
)

func init() {
	app := sco.GetApp()
	app.Options.Id = "gate_1"
	app.Options.Mid = 200
	app.Options.Frontend.WsAddr = "192.168.1.88:6000"
	app.Options.Cluster = true
	app.Options.RpcServer.Laddr = "192.168.1.88:8000"

	app.Options.Discovery.Endpoints = []string{
		"http://192.168.1.69:2379",
		"http://192.168.1.69:2479",
		"http://192.168.1.69:2579",
	}

	app.SetHandler(&Hander{})
}
