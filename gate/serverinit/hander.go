// /////////////////////////////////////////////////////////////////////////////
// 消息处理

package serverinit

import (
	"fmt"

	"github.com/zpab123/sco/network"
)

// /////////////////////////////////////////////////////////////////////////////
// 消息处理

type Hander struct {
}

func (this *Hander) OnPacket(a *network.Agent, pkt *network.Packet) {
	fmt.Println("OnPacket")
	if pkt.GetMid() == 123 {
		fmt.Println(pkt.ReadString())
	}
}
