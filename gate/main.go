// /////////////////////////////////////////////////////////////////////////////
// gate 服务器

package main

import (
	"github.com/zpab123/world" // world 库
)

// 主入口
func main() {
	// 创建 app
	app := world.CreateApp("gate")

	// 运行 app
	app.Run()
}
