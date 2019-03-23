package main

import (
	"encoding/binary"
	"fmt"
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

		if len(data) < 3 {
			break
		}

		go this.handlePacket(data)
	}

	return
}

// packet 层处理
func (this *Session) handlePacket(data []byte) {
	t := uint8(data[0])                         // 类型
	le := binary.LittleEndian.Uint16(data[1:3]) // 长度

	switch t {
	case protocol.C_TYPE_HANDSHAKE: // 客户端->服务器握手请求
	case protocol.C_TYPE_DATA: // 数据类
		this.handleMsg(le, data[3:])
	}
}

// message 层处理
func (this *Session) handleMsg(length uint16, body []byte) {
	if length < 4 {
		return
	}

	mtype := binary.LittleEndian.Uint16(body[0:2])
	stype := binary.LittleEndian.Uint16(body[2:4])

	if length > 4 {
		this.distribute(mtype, stype, body[4:])
	} else {
		this.distribute(mtype, stype, nil)
	}
}

// 分发消息
func (this *Session) distribute(mtype uint16, stype uint16, data []byte) {
	fmt.Println(mtype)
	fmt.Println(stype)
	fmt.Println(data)

	switch mtype {
	case protocol.C_MTYPE_LOGIN:
		{

		}
	default:
		panic("主协议错误")
	}
}
