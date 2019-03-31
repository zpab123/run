package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"run/light/protocol" // 消息协议
	"time"

	"github.com/gorilla/websocket" // websocket 库
)

// 会话对象
type Session struct {
	Connection   *websocket.Conn // websocket 连接对象
	Md5          string          // 房间加密信息
	playerID     uint32          // 玩家ID
	id           uint32          // sessionID
	lastRecvTime time.Time       // 上次接收到消息的时间
	lastSendTime time.Time       // 上次发送消息的时间
}

// 接收数据
func (this *Session) Recv() {
	for {
		_, data, err := this.Connection.ReadMessage()

		//data = bytes.TrimSpace(bytes.Replace(data, newline, space, -1))

		if nil != err {
			break
		}

		if len(data) < 4 {
			break
		}

		go this.handlePacket(data)
	}

	return
}

// packet 层处理
func (this *Session) handlePacket(data []byte) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常捕获", err)
		}
	}()

	mid := binary.LittleEndian.Uint16(data[0:2]) // 主id
	le := binary.LittleEndian.Uint16(data[2:4])  // 长度

	switch mid {
	case protocol.C_MID_HANDSHAKE: // 客户端->服务器握手请求
	case protocol.C_MID_HANDSHAKE_ACK: // 客户端->服务器握手ACK
	case protocol.C_MID_HEARTBEAT: // 心跳消息
	case protocol.C_MID_DATA: // 数据类
		this.handleMsg(mid, le, data[4:])
	default:
		panic("主协议错误")
	}
}

// message 层处理
func (this *Session) handleMsg(mid uint16, length uint16, body []byte) {
	if length < 2 {
		return
	}

	sid := binary.LittleEndian.Uint16(body[0:2])

	if length > 2 {
		this.distribute(mid, sid, body[2:])
	} else {
		this.distribute(mid, sid, nil)
	}
}

// 分发消息
func (this *Session) distribute(mid uint16, sid uint16, data []byte) {
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

	this.send(protocol.C_MID_DATA, 1, res)
}

// 发送数据
func (this *Session) send(mtype uint16, stype uint16, data interface{}) {
	b, err := json.Marshal(data)

	if nil != err {
		fmt.Println(err)
		return
	}

	// packet
	pkt := make([]byte, 2+2+2+len(b))
	binary.LittleEndian.PutUint16(pkt[0:2], protocol.C_MID_DATA) // 主id
	binary.LittleEndian.PutUint16(pkt[2:4], uint16(len(b)+2))    // 长度

	// body
	binary.LittleEndian.PutUint16(pkt[4:6], stype) // 子id
	copy(pkt[6:], b[:])

	fmt.Println(string(b))
	fmt.Println(pkt[6:])

	this.Connection.WriteMessage(websocket.BinaryMessage, pkt)
}

// 设置玩家id
func (this *Session) bindPlayer(playerID uint32) {
	this.playerID = playerID
}
