package main

import (
	"github.com/gorilla/websocket" // websocket 库
)

// 会话对象
type Session struct {
	Connection *websocket.Conn // websocket 连接对象
	Md5        string          // 房间加密信息
}

// 接收数据
func (this *Session) Recv() {
	return
}
