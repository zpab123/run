// /////////////////////////////////////////////////////////////////////////////
// 匹配服务

package main

import (
	"time"

	"github.com/zpab123/zaplog"
)

// /////////////////////////////////////////////////////////////////////////////
// 初始化

var (
	vs_1v1 []uint32 = make([]uint32, 0) // 1v1 列表
)

// /////////////////////////////////////////////////////////////////////////////
// match

// 1v1匹配服务
type match struct {
}

func newmatch() *match {
	m := match{}
	return &m
}

// 启动匹配服务
func (this *match) run() {
	// 每隔5秒匹配1次
	t := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-t.C:
			zaplog.Debug(
				"列表长度",
				zaplog.Uint32("ln", uint32(len(vs_1v1))),
			)
		}
	}
}
