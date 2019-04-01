package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"run/light/protocol" // 协议
	"time"

	"github.com/gorilla/websocket" // websocket 库
)

var (
	//raddr = "ws://192.168.0.222:8888/ws"
	raddr = "ws://127.0.0.1:8888/ws"
)

type robot struct {
	addr   string          // 服务器地址
	wsConn *websocket.Conn // 连接
	stop   chan struct{}   // 停止信号
}

// 新建机器人
func NewRobot() *robot {
	r := &robot{
		addr: "ws://192.168.0.222:8888/ws",
		stop: make(chan struct{}),
	}

	return r
}

// 启动
func (this *robot) Run() {
	this.runGame()
	this.login()
	go this.game()
	this.readData()
	//go this.readData()

	//<-this.stop

	fmt.Println("机器人退出")
}

// 启动游戏
func (this *robot) runGame() {
	conn, _, err := websocket.DefaultDialer.Dial(this.addr, nil)
	if nil != err {
		fmt.Println(err.Error())
	}

	this.wsConn = conn
}

// 登录
func (this *robot) login() {

}

// 游戏行为
func (this *robot) game() {
	n := rand.Int()%5 + 3
	t := time.Duration(n)

	for {
		select {
		case <-time.After(t * time.Second):
			this.sendDta()
		}
	}
}

// 数据接收
func (this *robot) readData() {
	defer func() {
		// this.stop <- struct{}{}
	}()

	if nil == this.wsConn {
		return
	}

	for {
		_, buf, err := this.wsConn.ReadMessage()

		if nil != err {
			break
		}

		if len(buf) < 4 {
			break
		}

		this.handlePacket(buf)
	}
}

// 结束游戏
func (this *robot) end() {
	this.wsConn.Close()
}

// packet 层处理
func (this *robot) handlePacket(data []byte) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常捕获", err)
		}
	}()

	mid := binary.LittleEndian.Uint16(data[0:2]) // 主id
	le := binary.LittleEndian.Uint16(data[2:4])  // 长度

	fmt.Println("mid", mid)

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
func (this *robot) handleMsg(mid uint16, length uint16, body []byte) {
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
func (this *robot) distribute(mid uint16, sid uint16, data []byte) {
	res := &protocol.LoginRes{}

	err := json.Unmarshal(data, res)

	if err != nil {
		fmt.Println(err.Error())

		return
	} else {
		fmt.Println("res-code:", res.Code)
	}
}

func (this *robot) sendDta() {
	req := &protocol.Data{
		Zp: "123gtt",
		Cd: "456fhth",
		Ef: "789hhfhhd",
	}

	this.send(protocol.C_MID_DATA, 1, req)
}

// 发送数据
func (this *robot) send(mid uint16, sid uint16, data interface{}) {
	if nil == this.wsConn {
		return
	}

	b, err := json.Marshal(data)

	if nil != err {
		fmt.Println(err.Error())

		return
	}

	// packet
	pkt := make([]byte, 2+2+2+len(b))
	binary.LittleEndian.PutUint16(pkt[0:2], mid)              // 主id
	binary.LittleEndian.PutUint16(pkt[2:4], uint16(len(b)+2)) // 长度

	// body
	binary.LittleEndian.PutUint16(pkt[4:6], sid) // 子id
	copy(pkt[6:], b[:])

	this.wsConn.WriteMessage(websocket.BinaryMessage, pkt)
}
