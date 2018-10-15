////////////////////////////////////////////////////////////////////////////////
// master 服务器

package main

import (
	"github.com/zpab123/scorpio"

	"go.uber.org/zap"
)

////////////////////////////////////////////////////////////////////////////////
// 主函数入口

// 主函数入口
func main() {
	// 日志变量
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 创建app
	app, scucess := scorpio.CreateApp()
	if scucess == false {
		logger.Error("创建 app 失败")
		return
	}

	// 配置pp （加载组件）

	// 启动app
	app.Start()

	// 异常捕获
}
