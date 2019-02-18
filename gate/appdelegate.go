// /////////////////////////////////////////////////////////////////////////////
// app代理

package main

import (
	"github.com/zpab123/world/session" // session
	"github.com/zpab123/zaplog"        // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// AppDelegate 对象

type AppDelegate struct {
	clientPacketQueue chan *session.Message // 客户端消息队列
	serverPacketQueue chan *session.Message // 服务器消息队列
}

// 新建1个 AppDelegate
func NewAppDelegate() *AppDelegate {
	// 创建
	ad := &AppDelegate{
		clientPacketQueue: make(chan *session.Message, 10000),
		serverPacketQueue: make(chan *session.Message, 10000),
	}

	return ad
}

// 运行 AppDelegate
func (this *AppDelegate) Run() {
	for {
		select {
		case msg := <-this.clientPacketQueue:
			// 处理客户端消息
			zaplog.Debugf(msg.GetPacket().GetId())
		case msg := <-this.serverPacketQueue:
			//
			zaplog.Debugf(msg.GetPacket().GetId())
		default:
		}
	}
}

// 收到1个新的客户端消息
func (this *AppDelegate) OnClientMessage(ses *session.FrontendSession, msg *session.Message) {
	// 加入客户端消息队列
	this.clientPacketQueue <- msg
}

// 收到1个新的服务器消息
func (this *AppDelegate) OnServerMessage(ses *session.BackendSession, msg *session.Message) {
	this.serverPacketQueue <- msg
}

// 处理客户端消息
func handleClientPacket() {

}

// 处理服务器消息
func handleServerPacket() {

}
