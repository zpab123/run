// /////////////////////////////////////////////////////////////////////////////
// 服务器初始化

package serverinit

import (
	"github.com/zpab123/sco"
)

func init() {
	app := sco.GetApp()
	app.Options.NetOpt.WsAddr = "192.168.1.88:5036"
	app.Options.Cluster = true
	app.Options.RpcOpt.Laddr = "192.168.1.88:6036"
	app.RegisterHandler(&Hander{})
}
