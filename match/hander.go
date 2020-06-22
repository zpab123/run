// /////////////////////////////////////////////////////////////////////////////
// 消息处理

package main

import (
	"fmt"

	"github.com/zpab123/sco"
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
	var res string

	rpc := sco.Call(201, data)
	if rpc == nil {
		res = "没收到rpc数据"
		pkt.AppendString(res)
		return true, pkt.Data()
	} else {
		pkt.AppendBytes(rpc)
		return true, pkt.Data()
	}
}
