// /////////////////////////////////////////////////////////////////////////////
// 消息处理

package main

import (
	"fmt"
	"math/rand"

	"github.com/zpab123/sco"
	"github.com/zpab123/sco/network"
	"github.com/zpab123/zaplog"
)

// /////////////////////////////////////////////////////////////////////////////
// 初始化

func init() {
	sco.GetApp().SetHandler(&Handler{})
}

// /////////////////////////////////////////////////////////////////////////////
// 消息处理

type Handler struct {
}

func (this *Handler) OnData(data []byte) (bool, []byte) {
	fmt.Println("OnPacket")
	fmt.Println(string(data))

	pkt := network.NewPacket(300)
	var res string = "match 收到数据"
	pkt.AppendString(res)

	r := rand.Uint32()
	zaplog.Debug(
		"随机值",
		zaplog.Uint32("r", r),
	)

	vs_1v1 = append(vs_1v1, r)

	return true, pkt.Data()
}
