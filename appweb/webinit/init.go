// /////////////////////////////////////////////////////////////////////////////
// 网站初始化设置
package webinit

import (
	_ "appweb/config"
	_ "appweb/routers"
)

// 系统初始化
func init() {
	logInit() // log 初始化
}
