// /////////////////////////////////////////////////////////////////////////////
// chat 服务器

package main

import (
	"github.com/zpab123/zpworld" // zpworld 库
)

// 主入口
func main() {
	// 注册场景实体
	zpworld.RegisterSpace(&Space{})

	// 注册 Account 实体
	zpworld.RegisterEntity("Account", &Account{})

	// 阻塞
	for {

	}
}
