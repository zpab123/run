// /////////////////////////////////////////////////////////////////////////////
// chat 空间

package main

import (
	"github.com/zpab123/zpworld/entity" // 实体库
)

type Space struct {
	entity.Space // 自定义的场景类型必须继承一个引擎所提供的entity.Space类型
}
