// /////////////////////////////////////////////////////////////////////////////
// app代理

package main

import (
	"github.com/zpab123/sco/app"     // 1个通用服务器
	"github.com/zpab123/sco/network" // 网络
	"github.com/zpab123/sco/session" // 会话库
)

// /////////////////////////////////////////////////////////////////////////////
// AppDelegate 对象

type AppDelegate struct {
}

// 新建1个 AppDelegate
func NewAppDelegate() *AppDelegate {
	// 创建
	ad := &AppDelegate{}

	return ad
}

// app 初始化
func (this *AppDelegate) Init(app *app.Application) {
}

// app 开始运行
func (this *AppDelegate) OnRun(app *app.Application) {

}

// app 停止运行
func (this *AppDelegate) OnStop(app *app.Application) {

}

// 收到1个新的客户端消息
func (this *AppDelegate) OnClientMessage(ses *session.ClientSession, packet *network.Packet) {

}

// 收到1个新的服务器消息
func (this *AppDelegate) OnServerMessage(ses *session.ServerSession, packet *network.Packet) {

}
