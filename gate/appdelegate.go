// /////////////////////////////////////////////////////////////////////////////
// app代理

package main

import (
	"github.com/zpab123/world/app"     // app 服务器
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // 会话库
	"github.com/zpab123/zaplog"        // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// AppDelegate 对象

type AppDelegate struct {
	app               *app.Application     // app 对象
	clientPacketQueue chan *network.Packet // 客户端消息队列
	serverPacketQueue chan *network.Packet // 服务器消息队列
}

// 新建1个 AppDelegate
func NewAppDelegate() *AppDelegate {
	// 创建
	ad := &AppDelegate{
		clientPacketQueue: make(chan *network.Packet, 10000),
		serverPacketQueue: make(chan *network.Packet, 10000),
	}

	return ad
}

// app 创建成功
func (this *AppDelegate) OnCreat(app *app.Application) {
	this.app = app
}

// app 初始化
func (this *AppDelegate) OnInit(app *app.Application) {
	// dispatcherClient 组件配置
	disClientOpt := app.GetComponentMgr().GetDisClientOpt()
	if nil != disClientOpt {
		disClientOpt.Enable = true
	}
}

// app 开始运行
func (this *AppDelegate) OnRun(app *app.Application) {
	for {
		select {
		case pkt := <-this.clientPacketQueue:
			// 处理客户端消息
			zaplog.Debugf("%d", pkt.GetId())
		case pkt := <-this.serverPacketQueue:
			//
			zaplog.Debugf("%d", pkt.GetId())
		}
	}
}

// app 停止运行
func (this *AppDelegate) OnStop(app *app.Application) {

}

// 收到1个新的客户端消息
func (this *AppDelegate) OnClientMessage(ses *session.ClientSession, packet *network.Packet) {
	// 加入客户端消息队列
	this.clientPacketQueue <- packet
}

// 收到1个新的服务器消息
func (this *AppDelegate) OnServerMessage(ses *session.ServerSession, packet *network.Packet) {
	this.serverPacketQueue <- packet
}

// 处理客户端消息
func (this *AppDelegate) handleClientPacket(pkt *network.Packet) {
	id := pkt.GetId()

	switch id {
	case 1:
	default:
		zaplog.Panicf("收到客户端非法消息。id=%d", id)
	}
}

// 处理服务器消息
func handleServerPacket() {

}
