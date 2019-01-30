// /////////////////////////////////////////////////////////////////////////////
// app代理

package main

import (
	"github.com/zpab123/world/session" // session
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

// session 消息管理
func (this *AppDelegate) OnNewMessage(msg *session.Message) {
	// 加入客户端消息队列
}

// 运行 AppDelegate
func (this *AppDelegate) run() {
	for {
		select {
		//case item := <-clientPacketQueue:
		//处理客户端消息
		}
	}
}
