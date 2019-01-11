// handler 注册中心

package main

import (
	"github.com/zpab123/world/ifs" // 全局接口
)

// 注册消息处理

type Handle struct {
	// 继承 或者 实现某个接口
}

// 消息处理
func (this *Handle) OnMessage(session ifs.ISession, data ifs.IMessage) {

}
