// /////////////////////////////////////////////////////////////////////////////
// go 空文件模板

package main

import (
	"github.com/zpab123/world/model" // 接口-常量-数据类型
)

type AppDelegate struct {
	clientPacketQueue chan interface{} // 客户端消息队列
}

// 客户端消息
func (this *AppDelegate) OnClientMsg(ses model.ISession, msg interface{}) {
	// 加入客户端消息队列
}

// 服务器消息
func (this *AppDelegate) OnServerMsg(ses model.ISession, msg interface{}) {
	// 加入服务器消息队列
}

// 运行 AppDelegate
func (this *AppDelegate) run() {
	for {
		select {
		case item := <-clientPacketQueue:
			//处理客户端消息
		}
	}
}
