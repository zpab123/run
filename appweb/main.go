package main

import (
	_ "run/appweb/config"
	_ "run/appweb/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs" // 日志库
)

func main() {
	// 设置日志输出
	logs.Async(10000) // 异步日志缓冲 chan 的大小
	logs.SetLogger(logs.AdapterFile, `{"filename":"log.log","level":7,"maxdays":10}`)

	beego.Run()
}
