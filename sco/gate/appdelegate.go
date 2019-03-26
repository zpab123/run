// /////////////////////////////////////////////////////////////////////////////
// app代理

package main

import (
	"github.com/zpab123/sco/app" // 1个通用服务器库
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

// app 创建成功
func (this *AppDelegate) OnCreat(app *app.Application) {
}

// app 初始化
func (this *AppDelegate) OnInit(app *app.Application) {
}

// app 开始运行
func (this *AppDelegate) OnRun(app *app.Application) {

}

// app 停止运行
func (this *AppDelegate) OnStop(app *app.Application) {

}
