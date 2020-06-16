// /////////////////////////////////////////////////////////////////////////////
// remote 消息

package serverinit

import (
	"github.com/zpab123/sco/network"
)

// /////////////////////////////////////////////////////////////////////////////
// 消息处理

type Remote struct {
}

func (this *Remote) OnData(data []byte) []byte {
	pkt := network.NewPacket(201)
	str := "rpc 收到数据"
	pkt.AppendString(str)

	return pkt.GetBody()
}
