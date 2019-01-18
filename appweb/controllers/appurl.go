package controllers

import (
	"github.com/astaxie/beego"
)

// 请求需要显示的网页地址
type AppUrl struct {
	beego.Controller
}

func (this *AppUrl) Get() {
	this.Ctx.WriteString("www.baid.com")
}
