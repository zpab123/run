// /////////////////////////////////////////////////////////////////////////////
// gate 服务器

package main

import (
	"github.com/zpab123/sco" // scorpio 游戏框架
)

// 主入口
func main() {
	// 创建代理
	ad := NewAppDelegate()

	// 创建 app
	app := sco.CreateApp("gate", ad)

	// 运行 app
	app.Run()
}
