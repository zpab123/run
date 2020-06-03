// /////////////////////////////////////////////////////////////////////////////
// 服务器初始化

package serverinit

import (
	"github.com/zpab123/sco"
)

func init() {
	app := sco.GetApp()
	app.Options.Name = "chat_1"
	app.Options.AppType = 2
	app.Options.ServiceId = 201
	app.Options.Cluster = true
	app.Options.RpcServer.Laddr = "192.168.1.88:7036"
	app.RegisterHandler(&Hander{})
	app.SetRemoteService(&Remote{})
}
