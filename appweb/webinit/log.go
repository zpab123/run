// /////////////////////////////////////////////////////////////////////////////
// 日志始化设置

package webinit

import (
	"github.com/astaxie/beego/logs"
)

// 日志初始化
func logInit() {
	logs.Async(10000) // 异步日志缓冲 chan 的大小
	logs.SetLogger(logs.AdapterFile, `{"filename":"log.log","level":7,"maxdays":10}`)
}
