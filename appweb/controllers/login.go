package controllers

import (
	"time"

	"github.com/astaxie/beego"      //  引擎
	"github.com/astaxie/beego/logs" // 日志库
)

// 请求需要显示的网页地址
type Login struct {
	beego.Controller
}

// get 方法
func (this *Login) Get() {
	// 记录ip 和 时间
	logs.Info("----------------------------登录开始--------------------------------------")
	logs.Info("ip=", this.Ctx.Request.RemoteAddr)
	logs.Info("time=", time.Now())
	logs.Info("----------------------------登录结束--------------------------------------")

	// 返回数据
	this.Ctx.WriteString("ok")
}
