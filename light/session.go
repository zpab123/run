package main

import (
	"run/light/protocol" // 消息协议

	"github.com/gorilla/websocket" // websocket 库
)

// 会话对象
type Session struct {
	Connection *websocket.Conn // websocket 连接对象
	Md5        string          // 房间加密信息
}

// 接收数据
func (this *Session) Recv() {
	for {
		_, data, err := this.Connection.ReadMessage()

		if nil != err {
			break
		}

		if 0 == len(data) {
			break
		}

		go this.decode(data)
	}

	return
}

// 消息解码
func (this *Session) decode(data []byte) {

}

// 分发消息
func (this *Session) distribute(mt int, sub int, data interface{}) {
	switch mt {
	case protocol.C_MT_DATA:
		{

		}
	default:
		panic("主协议错误")
	}
}
