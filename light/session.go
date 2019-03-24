package main

import (
	"encoding/binary"
	"encoding/json"
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

		//data = bytes.TrimSpace(bytes.Replace(data, newline, space, -1))

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
	sss := &protocol.Data{}

	err := json.Unmarshal(data, sss)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sss)
	}

	// 回复测试
	res := &protocol.LoginRes{
		Code:  123,
		Email: "shgshag",
	}

	this.send(1, 1, res)
}

// 发送数据
func (this *Session) send(mtype uint16, stype uint16, data interface{}) {
	b, err := json.Marshal(data)

	if nil != err {
		fmt.Println(err)
		return
	}

	// packet
	pkt := make([]byte, len(b)+4+3)
	pkt[0] = protocol.C_TYPE_DATA
	binary.LittleEndian.PutUint16(pkt[1:3], uint16(len(b)+4))

	// body
	binary.LittleEndian.PutUint16(pkt[3:5], mtype)
	binary.LittleEndian.PutUint16(pkt[5:7], stype)
	copy(pkt[7:], b[:])

	fmt.Println(string(b))
	fmt.Println(pkt[7:])

	this.Connection.WriteMessage(websocket.BinaryMessage, pkt)
}
