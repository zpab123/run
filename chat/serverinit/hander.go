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

func (this *Hander) OnData(data []byte) (bool, []byte) {
	fmt.Println("OnPacket")
	fmt.Println(string(data))

	pkt := network.NewPacket(201)
	pkt.AppendString("chat 收到数据")
	return true, pkt.Data()
}
