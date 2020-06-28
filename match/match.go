// /////////////////////////////////////////////////////////////////////////////
// 匹配服务

package main

import (
	"math/rand"
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
			this.soft()
		}
	}
}

// 从小到大排序
func (this *match) soft() {
	zaplog.Debug("-----s------")
	genRand()

	ln := len(vs_1v1)
	zaplog.Debug("长度", zaplog.Uint32("ln", uint32(ln)))
	var tem uint32
	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln-1; j++ {
			if vs_1v1[i] > vs_1v1[j] {
				tem = vs_1v1[i]
				vs_1v1[i] = vs_1v1[j]
				vs_1v1[j] = tem
			}
		}
	}

	zaplog.Debug("-----e------")
}

// 生成随机
func genRand() {
	vs_1v1 = vs_1v1[0:0]

	var r uint32
	for i := 0; i < 10000; i++ {
		r = rand.Uint32()
		vs_1v1 = append(vs_1v1, r)
	}
}
